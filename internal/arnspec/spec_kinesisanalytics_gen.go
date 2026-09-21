// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: kinesisanalytics
// Source: https://servicereference.us-east-1.amazonaws.com/v1/kinesisanalytics/kinesisanalytics.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "kinesisanalytics_application", Service: "kinesisanalytics", Resource: "application", Template: "arn:${Partition}:kinesisanalytics:${Region}:${Account}:application/${ApplicationName}"},
	})
}
