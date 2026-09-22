// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cloudwatch
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cloudwatch/cloudwatch.json
// Functions: 19
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cloudwatch_access_grant", Service: "cloudwatch", Resource: "access-grant", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:access-grant/${GrantId}"},
		{Name: "cloudwatch_access_profile", Service: "cloudwatch", Resource: "access-profile", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:access-profile/${ProfileId}"},
		{Name: "cloudwatch_alarm", Service: "cloudwatch", Resource: "alarm", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:alarm:${AlarmName}"},
		{Name: "cloudwatch_alarm_mute_rule", Service: "cloudwatch", Resource: "alarm-mute-rule", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:alarm-mute-rule:${AlarmMuteRuleName}"},
		{Name: "cloudwatch_alert", Service: "cloudwatch", Resource: "alert", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:alert/${AlertId}"},
		{Name: "cloudwatch_dashboard", Service: "cloudwatch", Resource: "dashboard", Template: "arn:${Partition}:cloudwatch::${Account}:dashboard/${DashboardName}"},
		{Name: "cloudwatch_dataset", Service: "cloudwatch", Resource: "dataset", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:dataset/${DatasetId}"},
		{Name: "cloudwatch_domain", Service: "cloudwatch", Resource: "domain", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:domain/${DomainId}"},
		{Name: "cloudwatch_ingestion_endpoint", Service: "cloudwatch", Resource: "ingestion-endpoint", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:ingestion-endpoint/${IngestionEndpointName}/${IngestionEndpointId}"},
		{Name: "cloudwatch_insight_rule", Service: "cloudwatch", Resource: "insight-rule", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:insight-rule/${InsightRuleName}"},
		{Name: "cloudwatch_integration", Service: "cloudwatch", Resource: "integration", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:integration/${IntegrationId}"},
		{Name: "cloudwatch_metric_stream", Service: "cloudwatch", Resource: "metric-stream", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:metric-stream/${MetricStreamName}"},
		{Name: "cloudwatch_omni_dashboard", Service: "cloudwatch", Resource: "omni-dashboard", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:omni-dashboard/${DashboardId}"},
		{Name: "cloudwatch_organization_access_grant", Service: "cloudwatch", Resource: "organization-access-grant", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:organization-access-grant/${GrantId}"},
		{Name: "cloudwatch_organization_domain", Service: "cloudwatch", Resource: "organization-domain", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:organization-domain/${DomainId}"},
		{Name: "cloudwatch_service", Service: "cloudwatch", Resource: "service", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:service/${ServiceName}-${UniqueAttributesHex}"},
		{Name: "cloudwatch_slo", Service: "cloudwatch", Resource: "slo", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:slo/${SloName}"},
		{Name: "cloudwatch_space", Service: "cloudwatch", Resource: "space", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:space/${SpaceId}"},
		{Name: "cloudwatch_view", Service: "cloudwatch", Resource: "view", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:view/${ViewName}"},
	})
}
