// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: qdeveloper
// Source: https://servicereference.us-east-1.amazonaws.com/v1/qdeveloper/qdeveloper.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "qdeveloper_code_transformation", Service: "qdeveloper", Resource: "codeTransformation", Template: "arn:${Partition}:qdeveloper:${Region}:${Account}:codeTransformation/${Identifier}"},
	})
}
