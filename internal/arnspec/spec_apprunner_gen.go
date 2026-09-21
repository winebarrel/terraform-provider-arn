// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: apprunner
// Source: https://servicereference.us-east-1.amazonaws.com/v1/apprunner/apprunner.json
// Functions: 7
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "apprunner_autoscalingconfiguration", Service: "apprunner", Resource: "autoscalingconfiguration", Template: "arn:${Partition}:apprunner:${Region}:${Account}:autoscalingconfiguration/${AutoscalingConfigurationName}/${AutoscalingConfigurationVersion}/${AutoscalingConfigurationId}"},
		{Name: "apprunner_connection", Service: "apprunner", Resource: "connection", Template: "arn:${Partition}:apprunner:${Region}:${Account}:connection/${ConnectionName}/${ConnectionId}"},
		{Name: "apprunner_observabilityconfiguration", Service: "apprunner", Resource: "observabilityconfiguration", Template: "arn:${Partition}:apprunner:${Region}:${Account}:observabilityconfiguration/${ObservabilityConfigurationName}/${ObservabilityConfigurationVersion}/${ObservabilityConfigurationId}"},
		{Name: "apprunner_service", Service: "apprunner", Resource: "service", Template: "arn:${Partition}:apprunner:${Region}:${Account}:service/${ServiceName}/${ServiceId}"},
		{Name: "apprunner_vpcconnector", Service: "apprunner", Resource: "vpcconnector", Template: "arn:${Partition}:apprunner:${Region}:${Account}:vpcconnector/${VpcConnectorName}/${VpcConnectorVersion}/${VpcConnectorId}"},
		{Name: "apprunner_vpcingressconnection", Service: "apprunner", Resource: "vpcingressconnection", Template: "arn:${Partition}:apprunner:${Region}:${Account}:vpcingressconnection/${VpcIngressConnectionName}/${VpcIngressConnectionId}"},
		{Name: "apprunner_webacl", Service: "apprunner", Resource: "webacl", Template: "arn:${Partition}:wafv2:${Region}:${Account}:${Scope}/webacl/${Name}/${Id}"},
	})
}
