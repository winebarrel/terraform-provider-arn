// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: kinesis
// Source: https://servicereference.us-east-1.amazonaws.com/v1/kinesis/kinesis.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "kinesis_channel", Service: "kinesis", Resource: "channel", Template: "arn:${Partition}:kinesis:${Region}:${Account}:channel/${ChannelId}"},
		{Name: "kinesis_consumer", Service: "kinesis", Resource: "consumer", Template: "arn:${Partition}:kinesis:${Region}:${Account}:${StreamType}/${StreamName}/consumer/${ConsumerName}:${ConsumerCreationTimpstamp}"},
		{Name: "kinesis_kms_key", Service: "kinesis", Resource: "kmsKey", Template: "arn:${Partition}:kms:${Region}:${Account}:key/${KeyId}"},
		{Name: "kinesis_stream", Service: "kinesis", Resource: "stream", Template: "arn:${Partition}:kinesis:${Region}:${Account}:stream/${StreamName}"},
	})
}
