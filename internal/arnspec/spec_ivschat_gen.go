// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: ivschat
// Source: https://servicereference.us-east-1.amazonaws.com/v1/ivschat/ivschat.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "ivschat_logging_configuration", Service: "ivschat", Resource: "Logging-Configuration", Template: "arn:${Partition}:ivschat:${Region}:${Account}:logging-configuration/${ResourceId}"},
		{Name: "ivschat_room", Service: "ivschat", Resource: "Room", Template: "arn:${Partition}:ivschat:${Region}:${Account}:room/${ResourceId}"},
	})
}
