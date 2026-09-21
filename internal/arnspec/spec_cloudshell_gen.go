// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cloudshell
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cloudshell/cloudshell.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cloudshell_environment", Service: "cloudshell", Resource: "Environment", Template: "arn:${Partition}:cloudshell:${Region}:${Account}:environment/${EnvironmentId}"},
	})
}
