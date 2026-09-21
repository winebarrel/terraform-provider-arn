// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: ecr
// Source: https://servicereference.us-east-1.amazonaws.com/v1/ecr/ecr.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "ecr_repository", Service: "ecr", Resource: "repository", Template: "arn:${Partition}:ecr:${Region}:${Account}:repository/${RepositoryName}"},
	})
}
