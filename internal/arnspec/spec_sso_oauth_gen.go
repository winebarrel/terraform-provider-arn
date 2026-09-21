// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: sso-oauth
// Source: https://servicereference.us-east-1.amazonaws.com/v1/sso-oauth/sso-oauth.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "sso_oauth_application", Service: "sso-oauth", Resource: "Application", Template: "arn:${Partition}:sso::${AccountId}:application/${InstanceId}/${ApplicationId}"},
	})
}
