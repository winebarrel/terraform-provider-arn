// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: ecr-public
// Source: https://servicereference.us-east-1.amazonaws.com/v1/ecr-public/ecr-public.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "ecr_public_registry", Service: "ecr-public", Resource: "registry", Template: "arn:${Partition}:ecr-public::${Account}:registry/${RegistryId}"},
		{Name: "ecr_public_repository", Service: "ecr-public", Resource: "repository", Template: "arn:${Partition}:ecr-public::${Account}:repository/${RepositoryName}"},
	})
}
