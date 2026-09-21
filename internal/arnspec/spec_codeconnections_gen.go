// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: codeconnections
// Source: https://servicereference.us-east-1.amazonaws.com/v1/codeconnections/codeconnections.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "codeconnections_connection", Service: "codeconnections", Resource: "Connection", Template: "arn:${Partition}:codeconnections:${Region}:${Account}:connection/${ConnectionId}"},
		{Name: "codeconnections_host", Service: "codeconnections", Resource: "Host", Template: "arn:${Partition}:codeconnections:${Region}:${Account}:host/${HostId}"},
		{Name: "codeconnections_repository_link", Service: "codeconnections", Resource: "RepositoryLink", Template: "arn:${Partition}:codeconnections:${Region}:${Account}:repository-link/${RepositoryLinkId}"},
	})
}
