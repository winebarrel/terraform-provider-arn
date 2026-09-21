// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: sso
// Source: https://servicereference.us-east-1.amazonaws.com/v1/sso/sso.json
// Functions: 6
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "sso_account", Service: "sso", Resource: "Account", Template: "arn:${Partition}:sso:::account/${AccountId}"},
		{Name: "sso_application", Service: "sso", Resource: "Application", Template: "arn:${Partition}:sso::${AccountId}:application/${InstanceId}/${ApplicationId}"},
		{Name: "sso_application_provider", Service: "sso", Resource: "ApplicationProvider", Template: "arn:${Partition}:sso::aws:applicationProvider/${ApplicationProviderId}"},
		{Name: "sso_instance", Service: "sso", Resource: "Instance", Template: "arn:${Partition}:sso:::instance/${InstanceId}"},
		{Name: "sso_permission_set", Service: "sso", Resource: "PermissionSet", Template: "arn:${Partition}:sso:::permissionSet/${InstanceId}/${PermissionSetId}"},
		{Name: "sso_trusted_token_issuer", Service: "sso", Resource: "TrustedTokenIssuer", Template: "arn:${Partition}:sso::${AccountId}:trustedTokenIssuer/${InstanceId}/${TrustedTokenIssuerId}"},
	})
}
