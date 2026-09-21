// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cognito-sync
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cognito-sync/cognito-sync.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cognito_sync_dataset", Service: "cognito-sync", Resource: "dataset", Template: "arn:${Partition}:cognito-sync:${Region}:${Account}:identitypool/${IdentityPoolId}/identity/${IdentityId}/dataset/${DatasetName}"},
		{Name: "cognito_sync_identity", Service: "cognito-sync", Resource: "identity", Template: "arn:${Partition}:cognito-sync:${Region}:${Account}:identitypool/${IdentityPoolId}/identity/${IdentityId}"},
		{Name: "cognito_sync_identitypool", Service: "cognito-sync", Resource: "identitypool", Template: "arn:${Partition}:cognito-sync:${Region}:${Account}:identitypool/${IdentityPoolId}"},
	})
}
