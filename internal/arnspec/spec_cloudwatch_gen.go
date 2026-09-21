// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cloudwatch
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cloudwatch/cloudwatch.json
// Functions: 8
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cloudwatch_alarm", Service: "cloudwatch", Resource: "alarm", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:alarm:${AlarmName}"},
		{Name: "cloudwatch_alarm_mute_rule", Service: "cloudwatch", Resource: "alarm-mute-rule", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:alarm-mute-rule:${AlarmMuteRuleName}"},
		{Name: "cloudwatch_dashboard", Service: "cloudwatch", Resource: "dashboard", Template: "arn:${Partition}:cloudwatch::${Account}:dashboard/${DashboardName}"},
		{Name: "cloudwatch_dataset", Service: "cloudwatch", Resource: "dataset", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:dataset/${DatasetId}"},
		{Name: "cloudwatch_insight_rule", Service: "cloudwatch", Resource: "insight-rule", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:insight-rule/${InsightRuleName}"},
		{Name: "cloudwatch_metric_stream", Service: "cloudwatch", Resource: "metric-stream", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:metric-stream/${MetricStreamName}"},
		{Name: "cloudwatch_service", Service: "cloudwatch", Resource: "service", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:service/${ServiceName}-${UniqueAttributesHex}"},
		{Name: "cloudwatch_slo", Service: "cloudwatch", Resource: "slo", Template: "arn:${Partition}:cloudwatch:${Region}:${Account}:slo/${SloName}"},
	})
}
