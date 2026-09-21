// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: iam
// Source: https://servicereference.us-east-1.amazonaws.com/v1/iam/iam.json
// Functions: 15
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "iam_access_report", Service: "iam", Resource: "access-report", Template: "arn:${Partition}:iam::${Account}:access-report/${EntityPath}"},
		{Name: "iam_assumed_role", Service: "iam", Resource: "assumed-role", Template: "arn:${Partition}:iam::${Account}:assumed-role/${RoleName}/${RoleSessionName}"},
		{Name: "iam_delegation_request", Service: "iam", Resource: "delegation-request", Template: "arn:${Partition}:iam::${Account}:delegation-request/${DelegationRequestId}"},
		{Name: "iam_federated_user", Service: "iam", Resource: "federated-user", Template: "arn:${Partition}:iam::${Account}:federated-user/${UserName}"},
		{Name: "iam_group", Service: "iam", Resource: "group", Template: "arn:${Partition}:iam::${Account}:group/${GroupNameWithPath}"},
		{Name: "iam_instance_profile", Service: "iam", Resource: "instance-profile", Template: "arn:${Partition}:iam::${Account}:instance-profile/${InstanceProfileNameWithPath}"},
		{Name: "iam_mfa", Service: "iam", Resource: "mfa", Template: "arn:${Partition}:iam::${Account}:mfa/${MfaTokenIdWithPath}"},
		{Name: "iam_oidc_provider", Service: "iam", Resource: "oidc-provider", Template: "arn:${Partition}:iam::${Account}:oidc-provider/${OidcProviderName}"},
		{Name: "iam_policy", Service: "iam", Resource: "policy", Template: "arn:${Partition}:iam::${Account}:policy/${PolicyNameWithPath}"},
		{Name: "iam_role", Service: "iam", Resource: "role", Template: "arn:${Partition}:iam::${Account}:role/${RoleNameWithPath}"},
		{Name: "iam_role_template", Service: "iam", Resource: "role-template", Template: "arn:${Partition}:iam::aws:role-template/${AWSServicePrincipal}/${RoleTemplateName}:${RoleTemplateMajorVersion}"},
		{Name: "iam_saml_provider", Service: "iam", Resource: "saml-provider", Template: "arn:${Partition}:iam::${Account}:saml-provider/${SamlProviderName}"},
		{Name: "iam_server_certificate", Service: "iam", Resource: "server-certificate", Template: "arn:${Partition}:iam::${Account}:server-certificate/${CertificateNameWithPath}"},
		{Name: "iam_sms_mfa", Service: "iam", Resource: "sms-mfa", Template: "arn:${Partition}:iam::${Account}:sms-mfa/${MfaTokenIdWithPath}"},
		{Name: "iam_user", Service: "iam", Resource: "user", Template: "arn:${Partition}:iam::${Account}:user/${UserNameWithPath}"},
	})
}
