// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: kinesisvideo
// Source: https://servicereference.us-east-1.amazonaws.com/v1/kinesisvideo/kinesisvideo.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "kinesisvideo_channel", Service: "kinesisvideo", Resource: "channel", Template: "arn:${Partition}:kinesisvideo:${Region}:${Account}:channel/${ChannelName}/${CreationTime}"},
		{Name: "kinesisvideo_stream", Service: "kinesisvideo", Resource: "stream", Template: "arn:${Partition}:kinesisvideo:${Region}:${Account}:stream/${StreamName}/${CreationTime}"},
	})
}
