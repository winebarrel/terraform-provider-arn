// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: mobiletargeting
// Source: https://servicereference.us-east-1.amazonaws.com/v1/mobiletargeting/mobiletargeting.json
// Functions: 29
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "mobiletargeting_app", Service: "mobiletargeting", Resource: "app", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}"},
		{Name: "mobiletargeting_application_metrics", Service: "mobiletargeting", Resource: "application-metrics", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/kpis/daterange/${KpiName}"},
		{Name: "mobiletargeting_apps", Service: "mobiletargeting", Resource: "apps", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/*"},
		{Name: "mobiletargeting_attribute", Service: "mobiletargeting", Resource: "attribute", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/attributes/${AttributeType}"},
		{Name: "mobiletargeting_campaign", Service: "mobiletargeting", Resource: "campaign", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/campaigns/${CampaignId}"},
		{Name: "mobiletargeting_campaign_metrics", Service: "mobiletargeting", Resource: "campaign-metrics", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/campaigns/${CampaignId}/kpis/daterange/${KpiName}"},
		{Name: "mobiletargeting_channel", Service: "mobiletargeting", Resource: "channel", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/channels/${ChannelType}"},
		{Name: "mobiletargeting_channels", Service: "mobiletargeting", Resource: "channels", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/channels"},
		{Name: "mobiletargeting_endpoint", Service: "mobiletargeting", Resource: "endpoint", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/endpoints/${EndpointId}"},
		{Name: "mobiletargeting_event_stream", Service: "mobiletargeting", Resource: "event-stream", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/eventstream"},
		{Name: "mobiletargeting_events", Service: "mobiletargeting", Resource: "events", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/events"},
		{Name: "mobiletargeting_export_job", Service: "mobiletargeting", Resource: "export-job", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/jobs/export/${JobId}"},
		{Name: "mobiletargeting_import_job", Service: "mobiletargeting", Resource: "import-job", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/jobs/import/${JobId}"},
		{Name: "mobiletargeting_journey", Service: "mobiletargeting", Resource: "journey", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/journeys/${JourneyId}"},
		{Name: "mobiletargeting_journey_execution_activity_metrics", Service: "mobiletargeting", Resource: "journey-execution-activity-metrics", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/journeys/${JourneyId}/activities/${JourneyActivityId}/execution-metrics"},
		{Name: "mobiletargeting_journey_execution_metrics", Service: "mobiletargeting", Resource: "journey-execution-metrics", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/journeys/${JourneyId}/execution-metrics"},
		{Name: "mobiletargeting_journey_metrics", Service: "mobiletargeting", Resource: "journey-metrics", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/journeys/${JourneyId}/kpis/daterange/${KpiName}"},
		{Name: "mobiletargeting_journeys", Service: "mobiletargeting", Resource: "journeys", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/journeys"},
		{Name: "mobiletargeting_messages", Service: "mobiletargeting", Resource: "messages", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/messages"},
		{Name: "mobiletargeting_otp", Service: "mobiletargeting", Resource: "otp", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/otp"},
		{Name: "mobiletargeting_phone_number_validate", Service: "mobiletargeting", Resource: "phone-number-validate", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:phone/number/validate"},
		{Name: "mobiletargeting_recommender", Service: "mobiletargeting", Resource: "recommender", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:recommenders/${RecommenderId}"},
		{Name: "mobiletargeting_recommenders", Service: "mobiletargeting", Resource: "recommenders", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:recommenders/*"},
		{Name: "mobiletargeting_reports", Service: "mobiletargeting", Resource: "reports", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:reports"},
		{Name: "mobiletargeting_segment", Service: "mobiletargeting", Resource: "segment", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/segments/${SegmentId}"},
		{Name: "mobiletargeting_template", Service: "mobiletargeting", Resource: "template", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:templates/${TemplateName}/${TemplateType}"},
		{Name: "mobiletargeting_templates", Service: "mobiletargeting", Resource: "templates", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:templates"},
		{Name: "mobiletargeting_user", Service: "mobiletargeting", Resource: "user", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/users/${UserId}"},
		{Name: "mobiletargeting_verify_otp", Service: "mobiletargeting", Resource: "verify-otp", Template: "arn:${Partition}:mobiletargeting:${Region}:${Account}:apps/${AppId}/verify-otp"},
	})
}
