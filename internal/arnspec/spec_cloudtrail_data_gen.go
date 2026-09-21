// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cloudtrail-data
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cloudtrail-data/cloudtrail-data.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cloudtrail_data_channel", Service: "cloudtrail-data", Resource: "channel", Template: "arn:${Partition}:cloudtrail:${Region}:${Account}:channel/${ChannelId}"},
	})
}
