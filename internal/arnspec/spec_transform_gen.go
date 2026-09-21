// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: transform
// Source: https://servicereference.us-east-1.amazonaws.com/v1/transform/transform.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "transform_connector", Service: "transform", Resource: "connector", Template: "arn:${Partition}:transform:${Region}:${Account}:connector/${WorkspaceId}/${ConnectorId}"},
		{Name: "transform_profile", Service: "transform", Resource: "profile", Template: "arn:${Partition}:transform:${Region}:${Account}:profile/${Identifier}"},
	})
}
