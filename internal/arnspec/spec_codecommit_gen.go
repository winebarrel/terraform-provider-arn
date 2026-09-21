// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: codecommit
// Source: https://servicereference.us-east-1.amazonaws.com/v1/codecommit/codecommit.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "codecommit_repository", Service: "codecommit", Resource: "repository", Template: "arn:${Partition}:codecommit:${Region}:${Account}:${RepositoryName}"},
	})
}
