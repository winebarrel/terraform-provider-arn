// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: events
// Source: https://servicereference.us-east-1.amazonaws.com/v1/events/events.json
// Functions: 15
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "events_alias", Service: "events", Resource: "alias", Template: "arn:${Partition}:kms:${Region}:${Account}:alias/${Alias}"},
		{Name: "events_api_destination", Service: "events", Resource: "api-destination", Template: "arn:${Partition}:events:${Region}:${Account}:api-destination/${ApiDestinationName}"},
		{Name: "events_archive", Service: "events", Resource: "archive", Template: "arn:${Partition}:events:${Region}:${Account}:archive/${ArchiveName}"},
		{Name: "events_connection", Service: "events", Resource: "connection", Template: "arn:${Partition}:events:${Region}:${Account}:connection/${ConnectionName}"},
		{Name: "events_create_snapshot", Service: "events", Resource: "create-snapshot", Template: "arn:${Partition}:events:${Region}:${Account}:target/create-snapshot"},
		{Name: "events_endpoint", Service: "events", Resource: "endpoint", Template: "arn:${Partition}:events:${Region}:${Account}:endpoint/${EndpointName}"},
		{Name: "events_event_bus", Service: "events", Resource: "event-bus", Template: "arn:${Partition}:events:${Region}:${Account}:event-bus/${EventBusName}"},
		{Name: "events_event_source", Service: "events", Resource: "event-source", Template: "arn:${Partition}:events:${Region}::event-source/${EventSourceName}"},
		{Name: "events_key", Service: "events", Resource: "key", Template: "arn:${Partition}:kms:${Region}:${Account}:key/${KeyId}"},
		{Name: "events_reboot_instance", Service: "events", Resource: "reboot-instance", Template: "arn:${Partition}:events:${Region}:${Account}:target/reboot-instance"},
		{Name: "events_replay", Service: "events", Resource: "replay", Template: "arn:${Partition}:events:${Region}:${Account}:replay/${ReplayName}"},
		{Name: "events_rule_on_custom_event_bus", Service: "events", Resource: "rule-on-custom-event-bus", Template: "arn:${Partition}:events:${Region}:${Account}:rule/${EventBusName}/${RuleName}"},
		{Name: "events_rule_on_default_event_bus", Service: "events", Resource: "rule-on-default-event-bus", Template: "arn:${Partition}:events:${Region}:${Account}:rule/${RuleName}"},
		{Name: "events_stop_instance", Service: "events", Resource: "stop-instance", Template: "arn:${Partition}:events:${Region}:${Account}:target/stop-instance"},
		{Name: "events_terminate_instance", Service: "events", Resource: "terminate-instance", Template: "arn:${Partition}:events:${Region}:${Account}:target/terminate-instance"},
	})
}
