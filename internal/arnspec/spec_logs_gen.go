// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: logs
// Source: https://servicereference.us-east-1.amazonaws.com/v1/logs/logs.json
// Functions: 9
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "logs_anomaly_detector", Service: "logs", Resource: "anomaly-detector", Template: "arn:${Partition}:logs:${Region}:${Account}:anomaly-detector:${DetectorId}"},
		{Name: "logs_delivery", Service: "logs", Resource: "delivery", Template: "arn:${Partition}:logs:${Region}:${Account}:delivery:${DeliveryName}"},
		{Name: "logs_delivery_destination", Service: "logs", Resource: "delivery-destination", Template: "arn:${Partition}:logs:${Region}:${Account}:delivery-destination:${DeliveryDestinationName}"},
		{Name: "logs_delivery_source", Service: "logs", Resource: "delivery-source", Template: "arn:${Partition}:logs:${Region}:${Account}:delivery-source:${DeliverySourceName}"},
		{Name: "logs_destination", Service: "logs", Resource: "destination", Template: "arn:${Partition}:logs:${Region}:${Account}:destination:${DestinationName}"},
		{Name: "logs_log_group", Service: "logs", Resource: "log-group", Template: "arn:${Partition}:logs:${Region}:${Account}:log-group:${LogGroupName}"},
		{Name: "logs_log_stream", Service: "logs", Resource: "log-stream", Template: "arn:${Partition}:logs:${Region}:${Account}:log-group:${LogGroupName}:log-stream:${LogStreamName}"},
		{Name: "logs_lookup_table", Service: "logs", Resource: "lookup-table", Template: "arn:${Partition}:logs:${Region}:${Account}:lookup-table:${LookupTableName}"},
		{Name: "logs_scheduled_query", Service: "logs", Resource: "scheduled-query", Template: "arn:${Partition}:logs:${Region}:${Account}:scheduled-query:${ScheduledQueryId}"},
	})
}
