package arnspec_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/terraform-provider-arn/internal/arnspec"
)

func parse(t *testing.T, tmpl string) *arnspec.Spec {
	t.Helper()
	s := &arnspec.Spec{Name: "test", Template: tmpl}
	require.NoError(t, arnspec.Parse(s))
	return s
}

func TestParseClassifiesPlaceholders(t *testing.T) {
	s := parse(t, "arn:${Partition}:iam::${Account}:role/${RoleNameWithPath}")
	assert.Equal(t, []string{"role_name_with_path"}, s.Args)
	assert.True(t, s.NeedsAccount)
	assert.False(t, s.NeedsRegion, "an IAM ARN has an empty region field")

	s = parse(t, "arn:${Partition}:sqs:${Region}:${Account}:${QueueName}")
	assert.Equal(t, []string{"queue_name"}, s.Args)
	assert.True(t, s.NeedsAccount)
	assert.True(t, s.NeedsRegion)

	s = parse(t, "arn:${Partition}:s3:::${BucketName}")
	assert.False(t, s.NeedsAccount)
	assert.False(t, s.NeedsRegion)
}

func TestParseKeepsArgumentOrder(t *testing.T) {
	s := parse(t, "arn:${Partition}:access-analyzer:${Region}:${Account}:analyzer/${AnalyzerName}/archive-rule/${RuleName}")
	assert.Equal(t, []string{"analyzer_name", "rule_name"}, s.Args)
	assert.Equal(t, "AnalyzerName", s.ArgRaw(0))
	assert.Equal(t, "RuleName", s.ArgRaw(1))
}

func TestParseRejectsUnterminatedPlaceholder(t *testing.T) {
	assert.Error(t, arnspec.Parse(&arnspec.Spec{Name: "test", Template: "arn:${Partition:iam"}))
}

func TestBuild(t *testing.T) {
	v := arnspec.Values{Partition: "aws", Region: "ap-northeast-1", AccountID: "111111111111"}

	s := parse(t, "arn:${Partition}:iam::${Account}:role/${RoleNameWithPath}")
	got, err := s.Build(v, []string{"my-role"})
	require.NoError(t, err)
	assert.Equal(t, "arn:aws:iam::111111111111:role/my-role", got)

	s = parse(t, "arn:${Partition}:access-analyzer:${Region}:${Account}:analyzer/${AnalyzerName}/archive-rule/${RuleName}")
	got, err = s.Build(v, []string{"an", "rule"})
	require.NoError(t, err)
	assert.Equal(t, "arn:aws:access-analyzer:ap-northeast-1:111111111111:analyzer/an/archive-rule/rule", got)

	// A template whose only placeholder is the partition still builds.
	s = parse(t, "arn:${Partition}:apigateway:${Region}::/account")
	got, err = s.Build(v, nil)
	require.NoError(t, err)
	assert.Equal(t, "arn:aws:apigateway:ap-northeast-1::/account", got)
}

func TestBuildReportsMissingValues(t *testing.T) {
	s := parse(t, "arn:${Partition}:sqs:${Region}:${Account}:${QueueName}")

	_, err := s.Build(arnspec.Values{Partition: "aws", AccountID: "1"}, []string{"q"})
	require.ErrorContains(t, err, "needs a region")

	_, err = s.Build(arnspec.Values{Partition: "aws", Region: "us-east-1"}, []string{"q"})
	require.ErrorContains(t, err, "needs an account id")

	_, err = s.Build(arnspec.Values{Partition: "aws", Region: "us-east-1", AccountID: "1"}, []string{""})
	require.ErrorContains(t, err, "is empty")

	_, err = s.Build(arnspec.Values{Partition: "aws", Region: "us-east-1", AccountID: "1"}, nil)
	require.ErrorContains(t, err, "takes 1 argument(s), got 0")
}

func TestBuildUsesTheGivenPartition(t *testing.T) {
	s := parse(t, "arn:${Partition}:s3:::${BucketName}")
	got, err := s.Build(arnspec.Values{Partition: "aws-cn"}, []string{"b"})
	require.NoError(t, err)
	assert.Equal(t, "arn:aws-cn:s3:::b", got)
}

func TestSnakeCase(t *testing.T) {
	for _, tt := range []struct{ in, want string }{
		{"role", "role"},
		{"certificate-authority", "certificate_authority"},
		{"IntegrationResponse", "integration_response"},
		{"gatewayRoute", "gateway_route"},
		{"TLSInspectionConfiguration", "tls_inspection_configuration"},
		{"MLInputChannel", "ml_input_channel"},
		{"ExportedAPI", "exported_api"},
		{"CIS Scan Configuration", "cis_scan_configuration"},
		{"loadbalancer/app/", "loadbalancer_app"},
		{"bot alias", "bot_alias"},
		{"inspector2", "inspector2"},
		{"Suitedefinition", "suitedefinition"},
	} {
		assert.Equal(t, tt.want, arnspec.SnakeCase(tt.in), tt.in)
	}
}

// The generated table is the provider's public surface, so a few well-known
// entries are pinned. A rename here is a breaking change for every
// configuration that calls the function.
func TestGeneratedSpecs(t *testing.T) {
	require.NotEmpty(t, arnspec.All())

	for _, tt := range []struct {
		name string
		tmpl string
		args []string
	}{
		{"iam_role", "arn:${Partition}:iam::${Account}:role/${RoleNameWithPath}", []string{"role_name_with_path"}},
		{"s3_bucket", "arn:${Partition}:s3:::${BucketName}", []string{"bucket_name"}},
		{"sqs_queue", "arn:${Partition}:sqs:${Region}:${Account}:${QueueName}", []string{"queue_name"}},
		{"lambda_function", "arn:${Partition}:lambda:${Region}:${Account}:function:${FunctionName}", []string{"function_name"}},
	} {
		s, ok := arnspec.Lookup(tt.name)
		require.True(t, ok, tt.name)
		assert.Equal(t, tt.tmpl, s.Template, tt.name)
		assert.Equal(t, tt.args, s.Args, tt.name)
	}
}

// Every generated name must be usable as a Terraform identifier, and unique.
func TestGeneratedNamesAreValidAndUnique(t *testing.T) {
	seen := map[string]bool{}
	for _, s := range arnspec.All() {
		assert.Regexp(t, `^[a-z][a-z0-9_]*$`, s.Name)
		assert.False(t, seen[s.Name], "duplicate %s", s.Name)
		seen[s.Name] = true

		for _, a := range s.Args {
			assert.Regexp(t, `^[a-z][a-z0-9_]*$`, a, s.Name)
		}
	}
}

// Every generated template must build once it is given values and arguments.
func TestGeneratedSpecsAllBuild(t *testing.T) {
	v := arnspec.Values{Partition: "aws", Region: "ap-northeast-1", AccountID: "111111111111"}
	for _, s := range arnspec.All() {
		args := make([]string, len(s.Args))
		for i := range args {
			args[i] = "x"
		}
		got, err := s.Build(v, args)
		require.NoError(t, err, s.Name)
		assert.True(t, strings.HasPrefix(got, "arn:aws:"), "%s: %s", s.Name, got)
		assert.NotContains(t, got, "${", "%s: %s", s.Name, got)
	}
}

// A handful of resource types publish more than one ARN format, and the
// numeric suffix is positional on the feed's ARNFormats order. If AWS ever
// reorders that list, apigateway_authorizer would quietly start building the
// REST API shape instead of the HTTP API one, which no other test would
// notice. Pinning every multi-format function turns that into a failure at
// regeneration time.
func TestGeneratedMultiFormatFunctions(t *testing.T) {
	want := map[string]string{
		"apigateway_authorizer":             "arn:${Partition}:apigateway:${Region}::/apis/${ApiId}/authorizers/${AuthorizerId}",
		"apigateway_authorizer_2":           "arn:${Partition}:apigateway:${Region}::/restapis/${RestApiId}/authorizers/${AuthorizerId}",
		"apigateway_authorizers":            "arn:${Partition}:apigateway:${Region}::/apis/${ApiId}/authorizers",
		"apigateway_authorizers_2":          "arn:${Partition}:apigateway:${Region}::/restapis/${RestApiId}/authorizers",
		"apigateway_deployment":             "arn:${Partition}:apigateway:${Region}::/apis/${ApiId}/deployments/${DeploymentId}",
		"apigateway_deployment_2":           "arn:${Partition}:apigateway:${Region}::/restapis/${RestApiId}/deployments/${DeploymentId}",
		"apigateway_deployments":            "arn:${Partition}:apigateway:${Region}::/apis/${ApiId}/deployments",
		"apigateway_deployments_2":          "arn:${Partition}:apigateway:${Region}::/restapis/${RestApiId}/deployments",
		"apigateway_integration":            "arn:${Partition}:apigateway:${Region}::/apis/${ApiId}/integrations/${IntegrationId}",
		"apigateway_integration_2":          "arn:${Partition}:apigateway:${Region}::/restapis/${RestApiId}/resources/${ResourceId}/methods/${HttpMethodType}/integration",
		"apigateway_integration_response":   "arn:${Partition}:apigateway:${Region}::/apis/${ApiId}/integrations/${IntegrationId}/integrationresponses/${IntegrationResponseId}",
		"apigateway_integration_response_2": "arn:${Partition}:apigateway:${Region}::/restapis/${RestApiId}/resources/${ResourceId}/methods/${HttpMethodType}/integration/responses/${StatusCode}",
		"apigateway_model":                  "arn:${Partition}:apigateway:${Region}::/apis/${ApiId}/models/${ModelId}",
		"apigateway_model_2":                "arn:${Partition}:apigateway:${Region}::/restapis/${RestApiId}/models/${ModelName}",
		"apigateway_models":                 "arn:${Partition}:apigateway:${Region}::/apis/${ApiId}/models",
		"apigateway_models_2":               "arn:${Partition}:apigateway:${Region}::/restapis/${RestApiId}/models",
		"apigateway_stage":                  "arn:${Partition}:apigateway:${Region}::/apis/${ApiId}/stages/${StageName}",
		"apigateway_stage_2":                "arn:${Partition}:apigateway:${Region}::/restapis/${RestApiId}/stages/${StageName}",
		"apigateway_stages":                 "arn:${Partition}:apigateway:${Region}::/apis/${ApiId}/stages",
		"apigateway_stages_2":               "arn:${Partition}:apigateway:${Region}::/restapis/${RestApiId}/stages",
		"greengrass_deployment":             "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/groups/${GroupId}/deployments/${DeploymentId}",
		"greengrass_deployment_2":           "arn:${Partition}:greengrass:${Region}:${Account}:deployments:${DeploymentId}",
		"lex_bot":                           "arn:${Partition}:lex:${Region}:${Account}:bot/${BotId}",
		"lex_bot_2":                         "arn:${Partition}:lex:${Region}:${Account}:bot:${BotName}",
		"lex_bot_alias":                     "arn:${Partition}:lex:${Region}:${Account}:bot-alias/${BotId}/${BotAliasId}",
		"lex_bot_alias_2":                   "arn:${Partition}:lex:${Region}:${Account}:bot:${BotName}:${BotAlias}",
	}

	// Collect what the generated table actually declares, so a new
	// multi-format resource type shows up as a failure too rather than
	// slipping in unpinned.
	type key struct{ service, resource string }

	got := map[string]string{}
	byResource := map[key]int{}
	for _, s := range arnspec.All() {
		byResource[key{s.Service, s.Resource}]++
	}
	for _, s := range arnspec.All() {
		if byResource[key{s.Service, s.Resource}] > 1 {
			got[s.Name] = s.Template
		}
	}
	assert.Equal(t, want, got)
}

// A colon in one of the five structural fields shifts every field after it.
// backup_recovery_point parameterises the service field.
func TestBuildRejectsColonInAStructuralArgument(t *testing.T) {
	s := parse(t, "arn:${Partition}:${Vendor}:${Region}:*:${ResourceType}:${RecoveryPointId}")
	v := arnspec.Values{Partition: "aws", Region: "us-east-1"}

	_, err := s.Build(v, []string{"backup:evil", "rt", "rp"})
	require.ErrorContains(t, err, "is an ARN field and cannot contain a colon")

	got, err := s.Build(v, []string{"backup", "rt", "rp"})
	require.NoError(t, err)
	assert.Equal(t, "arn:aws:backup:us-east-1:*:rt:rp", got)
}

// Inside the resource part a colon is just a character, and what it means is
// up to the service. S3 object keys may contain one.
func TestBuildAllowsColonInTheResourcePart(t *testing.T) {
	v := arnspec.Values{Partition: "aws", Region: "us-east-1", AccountID: "111111111111"}

	s := parse(t, "arn:${Partition}:s3:::${BucketName}/${ObjectName}")
	got, err := s.Build(v, []string{"my-bucket", "a:b/c.txt"})
	require.NoError(t, err)
	assert.Equal(t, "arn:aws:s3:::my-bucket/a:b/c.txt", got)

	s = parse(t, "arn:${Partition}:lambda:${Region}:${Account}:function:${FunctionName}")
	got, err = s.Build(v, []string{"my-func:PROD"})
	require.NoError(t, err)
	assert.Equal(t, "arn:aws:lambda:us-east-1:111111111111:function:my-func:PROD", got)
}

// Only one template still puts an argument in a structural field, so pin it
// along with the ordinary case.
func TestGeneratedStructuralArguments(t *testing.T) {
	structural := map[string][]string{}
	for _, s := range arnspec.All() {
		for i, a := range s.Args {
			if s.ArgIsStructural(i) {
				structural[s.Name] = append(structural[s.Name], a)
			}
		}
	}

	// backup_recovery_point parameterises the service field. Every other
	// structural field is filled from the configuration.
	assert.Equal(t, map[string][]string{
		"backup_recovery_point": {"vendor"},
	}, structural)

	s, ok := arnspec.Lookup("iam_role")
	require.True(t, ok)
	assert.False(t, s.ArgIsStructural(0))
}

// The account field is identified by position, so a template that spells it
// ${AccountId} takes the account from the configuration like any other.
func TestParseReadsTheAccountFieldByPosition(t *testing.T) {
	for _, tt := range []struct{ name, tmpl string }{
		{"chime_meeting", "arn:${Partition}:chime:${Region}:${AccountId}:meeting/${MeetingId}"},
		{"kafka_vpc_connection", "arn:${Partition}:kafka:${Region}:${VpcOwnerAccount}:vpc-connection/${ClusterOwnerAccount}/${ClusterName}/${Uuid}"},
		{"account_account_in_organization", "arn:${Partition}:account::${ManagementAccountId}:account/o-${OrganizationId}/${MemberAccountId}"},
	} {
		s, ok := arnspec.Lookup(tt.name)
		require.True(t, ok, tt.name)
		require.Equal(t, tt.tmpl, s.Template, tt.name)
		assert.True(t, s.NeedsAccount, tt.name)
		assert.NotContains(t, s.Args, "account_id", tt.name)
	}

	s, ok := arnspec.Lookup("chime_meeting")
	require.True(t, ok)
	assert.Equal(t, []string{"meeting_id"}, s.Args)

	got, err := s.Build(arnspec.Values{Partition: "aws", Region: "us-east-1", AccountID: "111111111111"}, []string{"m1"})
	require.NoError(t, err)
	assert.Equal(t, "arn:aws:chime:us-east-1:111111111111:meeting/m1", got)
}

// ${AccountId} in the resource part stays an argument, which is what keeps
// the two apart in organizations, where one template carries both.
func TestParseKeepsResourcePartAccountIDAsAnArgument(t *testing.T) {
	s, ok := arnspec.Lookup("organizations_account")
	require.True(t, ok)
	assert.Equal(t, "arn:${Partition}:organizations::${Account}:account/o-${OrganizationId}/${AccountId}", s.Template)
	assert.True(t, s.NeedsAccount)
	assert.Equal(t, []string{"organization_id", "account_id"}, s.Args)

	got, err := s.Build(arnspec.Values{Partition: "aws", AccountID: "111111111111"}, []string{"o-abc", "222222222222"})
	require.NoError(t, err)
	assert.Equal(t, "arn:aws:organizations::111111111111:account/o-o-abc/222222222222", got)
}

// A slash is not a field separator, and is part of plenty of legitimate
// names: IAM role paths, S3 object keys, CloudWatch log group names.
func TestBuildAllowsSlashInArgument(t *testing.T) {
	s := parse(t, "arn:${Partition}:iam::${Account}:role/${RoleNameWithPath}")
	got, err := s.Build(arnspec.Values{Partition: "aws", AccountID: "111111111111"}, []string{"path/to/my-role"})
	require.NoError(t, err)
	assert.Equal(t, "arn:aws:iam::111111111111:role/path/to/my-role", got)
}

// The whole classification rests on the correspondence between an ARN field
// and the names AWS puts in it, so pin it across every template rather than
// for a few named functions. A new spelling at the account field would
// otherwise set NeedsAccount, silently drop an argument, and pass every other
// test here.
func TestGeneratedPlaceholderPositions(t *testing.T) {
	// Field index counted in colons: 1 partition, 2 service, 3 region,
	// 4 account, 5 and beyond the resource part.
	want := map[int]map[string]bool{
		1: {"Partition": true},
		2: {"Vendor": true},
		3: {"Region": true},
		4: {
			"Account":             true,
			"AccountId":           true,
			"ManagementAccountId": true,
			"VpcOwnerAccount":     true,
		},
	}

	counts := map[int]map[string]int{1: {}, 2: {}, 3: {}, 4: {}}

	for _, s := range arnspec.All() {
		rest, field := s.Template, 0
		for {
			i := strings.Index(rest, "${")
			if i < 0 {
				break
			}
			j := strings.Index(rest[i:], "}")
			require.GreaterOrEqual(t, j, 0, s.Name)

			name := rest[i+2 : i+j]
			field += strings.Count(rest[:i], ":")
			if field < 5 {
				assert.True(t, want[field][name],
					"%s: unexpected placeholder ${%s} in ARN field %d: %s", s.Name, name, field, s.Template)
				counts[field][name]++
			}
			rest = rest[i+j+1:]
		}
	}

	// The counts themselves, so a template moving between fields shows up
	// even when the name is already known.
	assert.Equal(t, map[int]map[string]int{
		1: {"Partition": 2321},
		2: {"Vendor": 1},
		3: {"Region": 2095},
		4: {"Account": 2120, "AccountId": 22, "ManagementAccountId": 1, "VpcOwnerAccount": 1},
	}, counts)
}
