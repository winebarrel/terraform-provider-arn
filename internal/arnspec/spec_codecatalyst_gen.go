// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: codecatalyst
// Source: https://servicereference.us-east-1.amazonaws.com/v1/codecatalyst/codecatalyst.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "codecatalyst_connections", Service: "codecatalyst", Resource: "connections", Template: "arn:${Partition}:codecatalyst:${Region}:${Account}:/connections/${ConnectionId}"},
		{Name: "codecatalyst_identity_center_applications", Service: "codecatalyst", Resource: "identity-center-applications", Template: "arn:${Partition}:codecatalyst:${Region}:${Account}:/identity-center-applications/${IdentityCenterApplicationId}"},
		{Name: "codecatalyst_project", Service: "codecatalyst", Resource: "project", Template: "arn:${Partition}:codecatalyst:::space/${SpaceId}/project/${ProjectId}"},
		{Name: "codecatalyst_space", Service: "codecatalyst", Resource: "space", Template: "arn:${Partition}:codecatalyst:::space/${SpaceId}"},
	})
}
