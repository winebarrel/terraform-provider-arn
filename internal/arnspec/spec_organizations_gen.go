// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: organizations
// Source: https://servicereference.us-east-1.amazonaws.com/v1/organizations/organizations.json
// Functions: 9
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "organizations_account", Service: "organizations", Resource: "account", Template: "arn:${Partition}:organizations::${Account}:account/o-${OrganizationId}/${AccountId}"},
		{Name: "organizations_awspolicy", Service: "organizations", Resource: "awspolicy", Template: "arn:${Partition}:organizations::aws:policy/${PolicyType}/p-${PolicyId}"},
		{Name: "organizations_handshake", Service: "organizations", Resource: "handshake", Template: "arn:${Partition}:organizations::${Account}:handshake/o-${OrganizationId}/${HandshakeType}/h-${HandshakeId}"},
		{Name: "organizations_organization", Service: "organizations", Resource: "organization", Template: "arn:${Partition}:organizations::${Account}:organization/o-${OrganizationId}"},
		{Name: "organizations_organizationalunit", Service: "organizations", Resource: "organizationalunit", Template: "arn:${Partition}:organizations::${Account}:ou/o-${OrganizationId}/ou-${OrganizationalUnitId}"},
		{Name: "organizations_policy", Service: "organizations", Resource: "policy", Template: "arn:${Partition}:organizations::${Account}:policy/o-${OrganizationId}/${PolicyType}/p-${PolicyId}"},
		{Name: "organizations_resourcepolicy", Service: "organizations", Resource: "resourcepolicy", Template: "arn:${Partition}:organizations::${Account}:resourcepolicy/o-${OrganizationId}/rp-${ResourcePolicyId}"},
		{Name: "organizations_responsibilitytransfer", Service: "organizations", Resource: "responsibilitytransfer", Template: "arn:${Partition}:organizations::${Account}:transfer/o-${OrganizationId}/${TransferType}/${TransferDirection}/rt-${ResponsibilityTransferId}"},
		{Name: "organizations_root", Service: "organizations", Resource: "root", Template: "arn:${Partition}:organizations::${Account}:root/o-${OrganizationId}/r-${RootId}"},
	})
}
