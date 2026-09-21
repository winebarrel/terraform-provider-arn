// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cognito-idp
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cognito-idp/cognito-idp.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cognito_idp_userpool", Service: "cognito-idp", Resource: "userpool", Template: "arn:${Partition}:cognito-idp:${Region}:${Account}:userpool/${UserPoolId}"},
		{Name: "cognito_idp_webacl", Service: "cognito-idp", Resource: "webacl", Template: "arn:${Partition}:wafv2:${Region}:${Account}:${Scope}/webacl/${Name}/${Id}"},
	})
}
