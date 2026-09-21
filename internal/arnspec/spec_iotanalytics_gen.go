// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: iotanalytics
// Source: https://servicereference.us-east-1.amazonaws.com/v1/iotanalytics/iotanalytics.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "iotanalytics_channel", Service: "iotanalytics", Resource: "channel", Template: "arn:${Partition}:iotanalytics:${Region}:${Account}:channel/${ChannelName}"},
		{Name: "iotanalytics_dataset", Service: "iotanalytics", Resource: "dataset", Template: "arn:${Partition}:iotanalytics:${Region}:${Account}:dataset/${DatasetName}"},
		{Name: "iotanalytics_datastore", Service: "iotanalytics", Resource: "datastore", Template: "arn:${Partition}:iotanalytics:${Region}:${Account}:datastore/${DatastoreName}"},
		{Name: "iotanalytics_pipeline", Service: "iotanalytics", Resource: "pipeline", Template: "arn:${Partition}:iotanalytics:${Region}:${Account}:pipeline/${PipelineName}"},
	})
}
