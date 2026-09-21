// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cognito-identity
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cognito-identity/cognito-identity.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cognito_identity_identitypool", Service: "cognito-identity", Resource: "identitypool", Template: "arn:${Partition}:cognito-identity:${Region}:${Account}:identitypool/${IdentityPoolId}"},
	})
}
