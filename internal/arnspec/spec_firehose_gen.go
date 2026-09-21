// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: firehose
// Source: https://servicereference.us-east-1.amazonaws.com/v1/firehose/firehose.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "firehose_deliverystream", Service: "firehose", Resource: "deliverystream", Template: "arn:${Partition}:firehose:${Region}:${Account}:deliverystream/${DeliveryStreamName}"},
	})
}
