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
func TestBuildRejectsColonInAStructuralArgument(t *testing.T) {
	// chime spells the account field ${AccountId}, so it arrives as an
	// argument rather than from the configuration.
	s := parse(t, "arn:${Partition}:chime:${Region}:${AccountId}:meeting/${MeetingId}")
	v := arnspec.Values{Partition: "aws", Region: "us-east-1"}

	_, err := s.Build(v, []string{"111111111111:evil", "m1"})
	require.ErrorContains(t, err, "is an ARN field and cannot contain a colon")

	got, err := s.Build(v, []string{"111111111111", "m1"})
	require.NoError(t, err)
	assert.Equal(t, "arn:aws:chime:us-east-1:111111111111:meeting/m1", got)
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

// The classification is derived from the template, so pin both sides of it.
func TestGeneratedStructuralArguments(t *testing.T) {
	structural := map[string][]string{}
	for _, s := range arnspec.All() {
		for i, a := range s.Args {
			if s.ArgIsStructural(i) {
				structural[s.Name] = append(structural[s.Name], a)
			}
		}
	}

	// Every template that parameterises one of the first five fields.
	assert.Equal(t, map[string][]string{
		"account_account_in_organization":                {"management_account_id"},
		"backup_recovery_point":                          {"vendor"},
		"chime_app_instance":                             {"account_id"},
		"chime_app_instance_bot":                         {"account_id"},
		"chime_app_instance_user":                        {"account_id"},
		"chime_channel":                                  {"account_id"},
		"chime_channel_flow":                             {"account_id"},
		"chime_media_insights_pipeline_configuration":    {"account_id"},
		"chime_media_pipeline":                           {"account_id"},
		"chime_media_pipeline_kinesis_video_stream_pool": {"account_id"},
		"chime_meeting":                                  {"account_id"},
		"chime_sip_media_application":                    {"account_id"},
		"chime_voice_connector":                          {"account_id"},
		"chime_voice_profile":                            {"account_id"},
		"chime_voice_profile_domain":                     {"account_id"},
		"datasync_agent":                                 {"account_id"},
		"datasync_discoveryjob":                          {"account_id"},
		"datasync_location":                              {"account_id"},
		"datasync_storagesystem":                         {"account_id"},
		"datasync_task":                                  {"account_id"},
		"datasync_taskexecution":                         {"account_id"},
		"kafka_vpc_connection":                           {"vpc_owner_account"},
		"sso_application":                                {"account_id"},
		"sso_oauth_application":                          {"account_id"},
		"sso_trusted_token_issuer":                       {"account_id"},
	}, structural)

	// The common case is the other way round.
	s, ok := arnspec.Lookup("iam_role")
	require.True(t, ok)
	assert.False(t, s.ArgIsStructural(0))
}

// A slash is not a field separator, and is part of plenty of legitimate
// names: IAM role paths, S3 object keys, CloudWatch log group names.
func TestBuildAllowsSlashInArgument(t *testing.T) {
	s := parse(t, "arn:${Partition}:iam::${Account}:role/${RoleNameWithPath}")
	got, err := s.Build(arnspec.Values{Partition: "aws", AccountID: "111111111111"}, []string{"path/to/my-role"})
	require.NoError(t, err)
	assert.Equal(t, "arn:aws:iam::111111111111:role/path/to/my-role", got)
}
