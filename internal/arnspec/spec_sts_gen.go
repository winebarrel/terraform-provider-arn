// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: sts
// Source: https://servicereference.us-east-1.amazonaws.com/v1/sts/sts.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "sts_context_provider", Service: "sts", Resource: "context-provider", Template: "arn:${Partition}:iam::aws:contextProvider/${ContextProviderName}"},
		{Name: "sts_federated_user", Service: "sts", Resource: "federated-user", Template: "arn:${Partition}:sts::${Account}:federated-user/${FederatedUserName}"},
		{Name: "sts_role", Service: "sts", Resource: "role", Template: "arn:${Partition}:iam::${Account}:role/${RoleNameWithPath}"},
		{Name: "sts_root_user", Service: "sts", Resource: "root-user", Template: "arn:${Partition}:iam::${Account}:root"},
		{Name: "sts_self_session", Service: "sts", Resource: "self-session", Template: "arn:${Partition}:sts::${Account}:self"},
	})
}
