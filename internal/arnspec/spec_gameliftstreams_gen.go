// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: gameliftstreams
// Source: https://servicereference.us-east-1.amazonaws.com/v1/gameliftstreams/gameliftstreams.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "gameliftstreams_application", Service: "gameliftstreams", Resource: "application", Template: "arn:${Partition}:gameliftstreams:${Region}:${Account}:application/${ApplicationId}"},
		{Name: "gameliftstreams_stream_group", Service: "gameliftstreams", Resource: "stream group", Template: "arn:${Partition}:gameliftstreams:${Region}:${Account}:streamgroup/${StreamGroupId}"},
	})
}
