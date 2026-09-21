// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: codestar-connections
// Source: https://servicereference.us-east-1.amazonaws.com/v1/codestar-connections/codestar-connections.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "codestar_connections_connection", Service: "codestar-connections", Resource: "Connection", Template: "arn:${Partition}:codestar-connections:${Region}:${Account}:connection/${ConnectionId}"},
		{Name: "codestar_connections_host", Service: "codestar-connections", Resource: "Host", Template: "arn:${Partition}:codestar-connections:${Region}:${Account}:host/${HostId}"},
		{Name: "codestar_connections_repository_link", Service: "codestar-connections", Resource: "RepositoryLink", Template: "arn:${Partition}:codestar-connections:${Region}:${Account}:repository-link/${RepositoryLinkId}"},
	})
}
