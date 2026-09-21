// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: wafv2
// Source: https://servicereference.us-east-1.amazonaws.com/v1/wafv2/wafv2.json
// Functions: 13
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "wafv2_agentcore_gateway", Service: "wafv2", Resource: "agentcore-gateway", Template: "arn:${Partition}:bedrock-agentcore:${Region}:${Account}:gateway/${GatewayId}"},
		{Name: "wafv2_amplify_app", Service: "wafv2", Resource: "amplify-app", Template: "arn:${Partition}:amplify:${Region}:${Account}:apps/${AppId}"},
		{Name: "wafv2_apigateway", Service: "wafv2", Resource: "apigateway", Template: "arn:${Partition}:apigateway:${Region}::/restapis/${ApiId}/stages/${StageName}"},
		{Name: "wafv2_apprunner", Service: "wafv2", Resource: "apprunner", Template: "arn:${Partition}:apprunner:${Region}:${Account}:service/${ServiceName}/${ServiceId}"},
		{Name: "wafv2_appsync", Service: "wafv2", Resource: "appsync", Template: "arn:${Partition}:appsync:${Region}:${Account}:apis/${GraphQLAPIId}"},
		{Name: "wafv2_ipset", Service: "wafv2", Resource: "ipset", Template: "arn:${Partition}:wafv2:${Region}:${Account}:${Scope}/ipset/${Name}/${Id}"},
		{Name: "wafv2_loadbalancer_app", Service: "wafv2", Resource: "loadbalancer/app/", Template: "arn:${Partition}:elasticloadbalancing:${Region}:${Account}:loadbalancer/app/${LoadBalancerName}/${LoadBalancerId}"},
		{Name: "wafv2_managedruleset", Service: "wafv2", Resource: "managedruleset", Template: "arn:${Partition}:wafv2:${Region}:${Account}:${Scope}/managedruleset/${Name}/${Id}"},
		{Name: "wafv2_regexpatternset", Service: "wafv2", Resource: "regexpatternset", Template: "arn:${Partition}:wafv2:${Region}:${Account}:${Scope}/regexpatternset/${Name}/${Id}"},
		{Name: "wafv2_rulegroup", Service: "wafv2", Resource: "rulegroup", Template: "arn:${Partition}:wafv2:${Region}:${Account}:${Scope}/rulegroup/${Name}/${Id}"},
		{Name: "wafv2_userpool", Service: "wafv2", Resource: "userpool", Template: "arn:${Partition}:cognito-idp:${Region}:${Account}:userpool/${UserPoolId}"},
		{Name: "wafv2_verified_access_instance", Service: "wafv2", Resource: "verified-access-instance", Template: "arn:${Partition}:ec2:${Region}:${Account}:verified-access-instance/${VerifiedAccessInstanceId}"},
		{Name: "wafv2_webacl", Service: "wafv2", Resource: "webacl", Template: "arn:${Partition}:wafv2:${Region}:${Account}:${Scope}/webacl/${Name}/${Id}"},
	})
}
