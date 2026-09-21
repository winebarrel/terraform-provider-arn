// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: aws-external-anthropic
// Source: https://servicereference.us-east-1.amazonaws.com/v1/aws-external-anthropic/aws-external-anthropic.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "aws_external_anthropic_workspace", Service: "aws-external-anthropic", Resource: "workspace", Template: "arn:${Partition}:aws-external-anthropic:${Region}:${Account}:workspace/${ResourceId}"},
	})
}
