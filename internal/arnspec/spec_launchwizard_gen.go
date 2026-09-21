// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: launchwizard
// Source: https://servicereference.us-east-1.amazonaws.com/v1/launchwizard/launchwizard.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "launchwizard_deployment", Service: "launchwizard", Resource: "deployment", Template: "arn:${Partition}:launchwizard:${Region}:${Account}:deployment/${DeploymentId}"},
	})
}
