// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: workmail
// Source: https://servicereference.us-east-1.amazonaws.com/v1/workmail/workmail.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "workmail_organization", Service: "workmail", Resource: "organization", Template: "arn:${Partition}:workmail:${Region}:${Account}:organization/${ResourceId}"},
	})
}
