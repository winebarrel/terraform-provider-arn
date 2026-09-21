// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: translate
// Source: https://servicereference.us-east-1.amazonaws.com/v1/translate/translate.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "translate_parallel_data", Service: "translate", Resource: "parallel-data", Template: "arn:${Partition}:translate:${Region}:${Account}:parallel-data/${ResourceName}"},
		{Name: "translate_terminology", Service: "translate", Resource: "terminology", Template: "arn:${Partition}:translate:${Region}:${Account}:terminology/${ResourceName}"},
	})
}
