package provider_test

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/terraform-provider-arn/internal/arnconf"
	"github.com/winebarrel/terraform-provider-arn/internal/provider"
)

// A fresh provider per invocation, so a test that calls t.Setenv on
// ARN_CONFIG gets a provider whose configuration cache is still empty.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"arn": func() (tfprotov6.ProviderServer, error) {
		return providerserver.NewProtocol6WithError(provider.New("test")())()
	},
}

const testConfigFile = `
account_id = "111111111111"
region     = "ap-northeast-1"

account "prod" {
  account_id = "222222222222"
}

account "us" {
  account_id = "333333333333"
  region     = "us-east-1"
}
`

// useConfig points ARN_CONFIG at a temporary file for the duration of a test.
// The acceptance harness runs Terraform in its own working directory, so the
// default relative .arn.hcl lookup would find nothing.
func useConfig(t *testing.T, body string) {
	t.Helper()
	p := filepath.Join(t.TempDir(), ".arn.hcl")
	if err := os.WriteFile(p, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv(arnconf.EnvConfig, p)
}

func okStep(t *testing.T, config, expected string) {
	t.Helper()
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks:   []tfversion.TerraformVersionCheck{tfversion.SkipBelow(tfversion.Version1_8_0)},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: config, Check: resource.TestCheckOutput("test", expected)},
		},
	})
}

// errStep expects the step to fail with a diagnostic containing every phrase,
// in order.
//
// The phrases are plain text, not patterns. Terraform hard-wraps a diagnostic
// to the terminal width, and the wrap falls in a different place depending on
// how long the temporary directory in the message happens to be, so a phrase
// that reads as one line locally can arrive split across two in CI. Writing
// the patterns by hand meant every one of them had to guess where that break
// would land.
func errStep(t *testing.T, config string, phrases ...string) {
	t.Helper()
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks:   []tfversion.TerraformVersionCheck{tfversion.SkipBelow(tfversion.Version1_8_0)},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: config, ExpectError: wrapTolerant(phrases)},
		},
	})
}

// wrapTolerant turns plain phrases into one pattern: every run of spaces
// matches any whitespace, and the phrases may be separated by anything.
func wrapTolerant(phrases []string) *regexp.Regexp {
	parts := make([]string, 0, len(phrases))
	for _, p := range phrases {
		words := strings.Fields(p)
		for i, w := range words {
			words[i] = regexp.QuoteMeta(w)
		}
		parts = append(parts, strings.Join(words, `\s+`))
	}
	return regexp.MustCompile(`(?s)` + strings.Join(parts, ".*"))
}

func out(expr string) string {
	return "output \"test\" {\n  value = " + expr + "\n}\n"
}

func TestFunction_DefaultAccount(t *testing.T) {
	useConfig(t, testConfigFile)
	okStep(t, out(`provider::arn::iam_role("my-role")`), "arn:aws:iam::111111111111:role/my-role")
}

func TestFunction_RegionAndAccount(t *testing.T) {
	useConfig(t, testConfigFile)
	okStep(t, out(`provider::arn::sqs_queue("my-queue")`), "arn:aws:sqs:ap-northeast-1:111111111111:my-queue")
}

func TestFunction_NoAccountOrRegionNeeded(t *testing.T) {
	useConfig(t, testConfigFile)
	okStep(t, out(`provider::arn::s3_bucket("my-bucket")`), "arn:aws:s3:::my-bucket")
}

func TestFunction_NamedAccount(t *testing.T) {
	useConfig(t, testConfigFile)
	okStep(t, out(`provider::arn::iam_role("my-role", { account = "prod" })`), "arn:aws:iam::222222222222:role/my-role")
}

func TestFunction_NamedAccountInheritsDefaultRegion(t *testing.T) {
	useConfig(t, testConfigFile)
	okStep(t, out(`provider::arn::sqs_queue("q", { account = "prod" })`), "arn:aws:sqs:ap-northeast-1:222222222222:q")
}

func TestFunction_NamedAccountOwnRegion(t *testing.T) {
	useConfig(t, testConfigFile)
	okStep(t, out(`provider::arn::sqs_queue("q", { account = "us" })`), "arn:aws:sqs:us-east-1:333333333333:q")
}

func TestFunction_OptionOverrides(t *testing.T) {
	useConfig(t, testConfigFile)
	okStep(t, out(`provider::arn::sqs_queue("q", { region = "eu-west-1" })`), "arn:aws:sqs:eu-west-1:111111111111:q")
	okStep(t, out(`provider::arn::iam_role("r", { account_id = "999999999999" })`), "arn:aws:iam::999999999999:role/r")
	okStep(t, out(`provider::arn::s3_bucket("b", { partition = "aws-cn" })`), "arn:aws-cn:s3:::b")
}

func TestFunction_MultipleArguments(t *testing.T) {
	useConfig(t, testConfigFile)
	okStep(t,
		out(`provider::arn::access_analyzer_archive_rule("an", "rule")`),
		"arn:aws:access-analyzer:ap-northeast-1:111111111111:analyzer/an/archive-rule/rule")
}

// The numeric suffix marks a second ARN format on the same resource type,
// which for apigateway is the older REST API shape.
func TestFunction_SecondARNFormat(t *testing.T) {
	useConfig(t, testConfigFile)
	okStep(t, out(`provider::arn::apigateway_authorizer("api1", "auth1")`),
		"arn:aws:apigateway:ap-northeast-1::/apis/api1/authorizers/auth1")
	okStep(t, out(`provider::arn::apigateway_authorizer_2("rest1", "auth1")`),
		"arn:aws:apigateway:ap-northeast-1::/restapis/rest1/authorizers/auth1")
}

func TestFunction_UnknownAccount(t *testing.T) {
	useConfig(t, testConfigFile)
	errStep(t, out(`provider::arn::iam_role("r", { account = "nope" })`), `unknown account "nope"`, `prod, us`)
}

func TestFunction_UnknownOption(t *testing.T) {
	useConfig(t, testConfigFile)
	errStep(t, out(`provider::arn::iam_role("r", { acccount = "prod" })`), `unknown option(s) acccount`)
}

func TestFunction_AccountAndAccountIDAreExclusive(t *testing.T) {
	useConfig(t, testConfigFile)
	errStep(t, out(`provider::arn::iam_role("r", { account = "prod", account_id = "9" })`), `mutually exclusive`)
}

func TestFunction_AtMostOneOptionsMap(t *testing.T) {
	useConfig(t, testConfigFile)
	errStep(t, out(`provider::arn::iam_role("r", {}, {})`), `at most one options map, got 2`)
}

func TestFunction_EmptyArgument(t *testing.T) {
	useConfig(t, testConfigFile)
	errStep(t, out(`provider::arn::iam_role("")`), `role_name_with_path (RoleNameWithPath) is empty`)
}

func TestFunction_MissingAccountID(t *testing.T) {
	useConfig(t, `region = "ap-northeast-1"`)
	errStep(t, out(`provider::arn::iam_role("r")`), `needs an account id`)
}

func TestFunction_MissingRegion(t *testing.T) {
	useConfig(t, `account_id = "111111111111"`)
	errStep(t, out(`provider::arn::sqs_queue("q")`), `needs a region`)
}

// The configuration file is required, even for an ARN that would need
// nothing from it.
func TestFunction_NoConfigFile(t *testing.T) {
	t.Setenv(arnconf.EnvConfig, filepath.Join(t.TempDir(), "absent.hcl"))
	errStep(t, out(`provider::arn::s3_bucket("b")`), `absent.hcl not found`)
	errStep(t, out(`provider::arn::iam_role("r")`), `absent.hcl not found`)
}

// The file has to exist, but it does not have to say anything.
func TestFunction_EmptyConfigFile(t *testing.T) {
	useConfig(t, "")
	okStep(t, out(`provider::arn::s3_bucket("b")`), "arn:aws:s3:::b")
	errStep(t, out(`provider::arn::iam_role("r")`), `needs an account id`)
}

func TestFunction_MalformedConfigFile(t *testing.T) {
	useConfig(t, `account_id = `)
	errStep(t, out(`provider::arn::iam_role("r")`), `Missing expression`)
}

func TestFunction_RejectsMalformedValues(t *testing.T) {
	useConfig(t, testConfigFile)

	errStep(t, out(`provider::arn::iam_role("r", { account_id = "abc" })`), `invalid account id "abc"`)
	errStep(t, out(`provider::arn::iam_role("r", { account_id = "12345" })`), `invalid account id`)
	errStep(t, out(`provider::arn::sqs_queue("q", { region = "not-a-region" })`), `invalid region "not-a-region"`)
	errStep(t, out(`provider::arn::s3_bucket("b", { partition = "nonsense" })`), `invalid partition "nonsense"`)
	errStep(t, out(`provider::arn::iam_role("r", { account_id = "111111111111:evil" })`), `invalid account id`)
}

// A colon in one of the ARN's five structural fields shifts every field after
// it. chime spells the account field ${AccountId}, so it arrives as an
// argument rather than from the configuration.
func TestFunction_RejectsColonInAStructuralArgument(t *testing.T) {
	useConfig(t, testConfigFile)
	errStep(t, out(`provider::arn::chime_meeting("111111111111:evil", "m1")`),
		`is an ARN field and cannot contain a colon`)
	okStep(t, out(`provider::arn::chime_meeting("111111111111", "m1")`),
		"arn:aws:chime:ap-northeast-1:111111111111:meeting/m1")
}

// Inside the resource part a colon is just a character. S3 object keys may
// contain one, and a lambda alias ARN is written with one.
func TestFunction_AllowsColonInTheResourcePart(t *testing.T) {
	useConfig(t, testConfigFile)
	okStep(t, out(`provider::arn::s3_object("my-bucket", "a:b/c.txt")`), "arn:aws:s3:::my-bucket/a:b/c.txt")
	okStep(t, out(`provider::arn::lambda_function_alias("fn", "PROD")`),
		"arn:aws:lambda:ap-northeast-1:111111111111:function:fn:PROD")
}

// The single most commonly hand-written IAM ARN: an AWS-managed policy.
func TestFunction_AWSManagedPolicy(t *testing.T) {
	useConfig(t, testConfigFile)
	okStep(t, out(`provider::arn::iam_policy("AdministratorAccess", { account_id = "aws" })`),
		"arn:aws:iam::aws:policy/AdministratorAccess")
}

// An ARN written for an IAM policy may wildcard any of the structural fields.
func TestFunction_Wildcards(t *testing.T) {
	useConfig(t, testConfigFile)
	okStep(t, out(`provider::arn::ec2_vpc("*", { region = "*" })`), "arn:aws:ec2:*:111111111111:vpc/*")
	okStep(t, out(`provider::arn::iam_role("*", { account_id = "*" })`), "arn:aws:iam::*:role/*")
	okStep(t, out(`provider::arn::s3_object("my-bucket", "*", { partition = "*" })`), "arn:*:s3:::my-bucket/*")
}

// chime spells the account field as a template placeholder, so it arrives as
// an argument. It gets the same check as { account_id = ... } does.
func TestFunction_ValidatesAnAccountFieldArgument(t *testing.T) {
	useConfig(t, testConfigFile)
	errStep(t, out(`provider::arn::chime_meeting("abc", "m1")`), `invalid account id "abc"`)
	okStep(t, out(`provider::arn::chime_meeting("222222222222", "m1")`),
		"arn:aws:chime:ap-northeast-1:222222222222:meeting/m1")
}

// A slash is part of plenty of legitimate names.
func TestFunction_AllowsSlashInArgument(t *testing.T) {
	useConfig(t, testConfigFile)
	okStep(t, out(`provider::arn::iam_role("path/to/my-role")`), "arn:aws:iam::111111111111:role/path/to/my-role")
}

// A bad value in the configuration file is reported even when the call site
// overrides nothing.
func TestFunction_RejectsMalformedConfigValues(t *testing.T) {
	useConfig(t, `
account_id = "111111111111"
region     = "nihon"
`)
	errStep(t, out(`provider::arn::sqs_queue("q")`), `invalid region "nihon"`)
}

// The text below is the diagnostic that failed in CI while passing locally:
// the temp directory path pushed "Missing expression" across a line break.
func TestWrapTolerant(t *testing.T) {
	const ciOutput = `Call to function "provider::arn::iam_role" failed:
/tmp/TestFunction_MalformedConfigFile2303570173/001/.arn.hcl:1,14-14: Missing
expression; Expected the start of an expression, but found the end of the
file..`

	// What the hand-written pattern did, and why it only failed in CI.
	assert.NotRegexp(t, `(?s)Missing expression`, ciOutput)
	assert.Regexp(t, wrapTolerant([]string{"Missing expression"}), ciOutput)
	assert.Regexp(t, wrapTolerant([]string{"found the end of the file"}), ciOutput)

	// Phrases must match in order, and metacharacters are literal.
	const wrapped = "unknown account \"nope\": declared\naccounts are prod, us."
	assert.Regexp(t, wrapTolerant([]string{`unknown account "nope"`, "prod, us"}), wrapped)
	assert.NotRegexp(t, wrapTolerant([]string{"prod, us", `unknown account "nope"`}), wrapped)
	assert.NotRegexp(t, wrapTolerant([]string{"unknown account .nope."}), wrapped)
}

// useConfigTree writes several files and points ARN_CONFIG at the first.
func useConfigTree(t *testing.T, root string, files map[string]string) {
	t.Helper()
	dir := t.TempDir()
	for name, body := range files {
		p := filepath.Join(dir, name)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	t.Setenv(arnconf.EnvConfig, filepath.Join(dir, root))
}

func TestFunction_Import(t *testing.T) {
	useConfigTree(t, ".arn.hcl", map[string]string{
		"common.hcl": `
account_id = "111111111111"
region     = "ap-northeast-1"

account "prod" {
  account_id = "222222222222"
}
`,
		".arn.hcl": `
import = "common.hcl"
`,
	})

	okStep(t, out(`provider::arn::iam_role("r")`), "arn:aws:iam::111111111111:role/r")
	okStep(t, out(`provider::arn::iam_role("r", { account = "prod" })`), "arn:aws:iam::222222222222:role/r")
	okStep(t, out(`provider::arn::sqs_queue("q", { account = "prod" })`), "arn:aws:sqs:ap-northeast-1:222222222222:q")
}

func TestFunction_ImportIsOverriddenByTheImportingFile(t *testing.T) {
	useConfigTree(t, ".arn.hcl", map[string]string{
		"common.hcl": `
account_id = "111111111111"
region     = "ap-northeast-1"
`,
		".arn.hcl": `
import = "common.hcl"
region = "us-east-1"
`,
	})

	okStep(t, out(`provider::arn::sqs_queue("q")`), "arn:aws:sqs:us-east-1:111111111111:q")
}

func TestFunction_ImportNotFound(t *testing.T) {
	useConfigTree(t, ".arn.hcl", map[string]string{
		".arn.hcl": `import = "missing.hcl"`,
	})
	errStep(t, out(`provider::arn::iam_role("r")`), `import "missing.hcl" not found`)
}
