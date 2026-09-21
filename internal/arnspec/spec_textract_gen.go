// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: textract
// Source: https://servicereference.us-east-1.amazonaws.com/v1/textract/textract.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "textract_adapter", Service: "textract", Resource: "adapter", Template: "arn:${Partition}:textract:${Region}:${Account}:/adapters/${AdapterId}"},
		{Name: "textract_adapterversion", Service: "textract", Resource: "adapterversion", Template: "arn:${Partition}:textract:${Region}:${Account}:/adapters/${AdapterId}/versions/${AdapterVersion}"},
	})
}
