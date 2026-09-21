package provider_test

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
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

func errStep(t *testing.T, config, errPattern string) {
	t.Helper()
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks:   []tfversion.TerraformVersionCheck{tfversion.SkipBelow(tfversion.Version1_8_0)},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: config, ExpectError: regexp.MustCompile(errPattern)},
		},
	})
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
	errStep(t, out(`provider::arn::iam_role("r", { account = "nope" })`), `(?s)unknown account "nope".*prod,\s+us`)
}

func TestFunction_UnknownOption(t *testing.T) {
	useConfig(t, testConfigFile)
	errStep(t, out(`provider::arn::iam_role("r", { acccount = "prod" })`), `unknown option\(s\) acccount`)
}

func TestFunction_AccountAndAccountIDAreExclusive(t *testing.T) {
	useConfig(t, testConfigFile)
	errStep(t, out(`provider::arn::iam_role("r", { account = "prod", account_id = "9" })`), `(?s)mutually\s+exclusive`)
}

func TestFunction_AtMostOneOptionsMap(t *testing.T) {
	useConfig(t, testConfigFile)
	errStep(t, out(`provider::arn::iam_role("r", {}, {})`), `(?s)at most one options\s+map, got 2`)
}

func TestFunction_EmptyArgument(t *testing.T) {
	useConfig(t, testConfigFile)
	errStep(t, out(`provider::arn::iam_role("")`), `role_name_with_path \(RoleNameWithPath\) is empty`)
}

func TestFunction_MissingAccountID(t *testing.T) {
	useConfig(t, `region = "ap-northeast-1"`)
	errStep(t, out(`provider::arn::iam_role("r")`), `(?s)needs an account\s+id`)
}

func TestFunction_MissingRegion(t *testing.T) {
	useConfig(t, `account_id = "111111111111"`)
	errStep(t, out(`provider::arn::sqs_queue("q")`), `(?s)needs a\s+region`)
}

// The configuration file is required, even for an ARN that would need
// nothing from it.
func TestFunction_NoConfigFile(t *testing.T) {
	t.Setenv(arnconf.EnvConfig, filepath.Join(t.TempDir(), "absent.hcl"))
	errStep(t, out(`provider::arn::s3_bucket("b")`), `(?s)absent\.hcl\s+not\s+found`)
	errStep(t, out(`provider::arn::iam_role("r")`), `(?s)absent\.hcl\s+not\s+found`)
}

// The file has to exist, but it does not have to say anything.
func TestFunction_EmptyConfigFile(t *testing.T) {
	useConfig(t, "")
	okStep(t, out(`provider::arn::s3_bucket("b")`), "arn:aws:s3:::b")
	errStep(t, out(`provider::arn::iam_role("r")`), `(?s)needs an account\s+id`)
}

func TestFunction_MalformedConfigFile(t *testing.T) {
	useConfig(t, `account_id = `)
	errStep(t, out(`provider::arn::iam_role("r")`), `(?s)Missing expression`)
}
