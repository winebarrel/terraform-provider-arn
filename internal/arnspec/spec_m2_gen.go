// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: m2
// Source: https://servicereference.us-east-1.amazonaws.com/v1/m2/m2.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "m2_application", Service: "m2", Resource: "Application", Template: "arn:${Partition}:m2:${Region}:${Account}:app/${ApplicationId}"},
		{Name: "m2_environment", Service: "m2", Resource: "Environment", Template: "arn:${Partition}:m2:${Region}:${Account}:env/${EnvironmentId}"},
	})
}
