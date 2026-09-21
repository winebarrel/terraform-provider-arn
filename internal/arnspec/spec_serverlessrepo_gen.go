// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: serverlessrepo
// Source: https://servicereference.us-east-1.amazonaws.com/v1/serverlessrepo/serverlessrepo.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "serverlessrepo_applications", Service: "serverlessrepo", Resource: "applications", Template: "arn:${Partition}:serverlessrepo:${Region}:${Account}:applications/${ResourceId}"},
	})
}
