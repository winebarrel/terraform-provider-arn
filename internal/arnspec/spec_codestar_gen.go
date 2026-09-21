// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: codestar
// Source: https://servicereference.us-east-1.amazonaws.com/v1/codestar/codestar.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "codestar_project", Service: "codestar", Resource: "project", Template: "arn:${Partition}:codestar:${Region}:${Account}:project/${ProjectId}"},
		{Name: "codestar_user", Service: "codestar", Resource: "user", Template: "arn:${Partition}:iam::${Account}:user/${AwsUserName}"},
	})
}
