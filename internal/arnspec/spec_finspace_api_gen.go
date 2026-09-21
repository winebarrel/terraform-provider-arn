// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: finspace-api
// Source: https://servicereference.us-east-1.amazonaws.com/v1/finspace-api/finspace-api.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "finspace_api_credential", Service: "finspace-api", Resource: "credential", Template: "arn:${Partition}:finspace-api:${Region}:${Account}:/credentials/programmatic"},
	})
}
