// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: eks
// Source: https://servicereference.us-east-1.amazonaws.com/v1/eks/eks.json
// Functions: 11
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "eks_access_entry", Service: "eks", Resource: "access-entry", Template: "arn:${Partition}:eks:${Region}:${Account}:access-entry/${ClusterName}/${IamIdentityType}/${IamIdentityAccountID}/${IamIdentityName}/${UUID}"},
		{Name: "eks_access_policy", Service: "eks", Resource: "access-policy", Template: "arn:${Partition}:eks::aws:cluster-access-policy/${AccessPolicyName}"},
		{Name: "eks_addon", Service: "eks", Resource: "addon", Template: "arn:${Partition}:eks:${Region}:${Account}:addon/${ClusterName}/${AddonName}/${UUID}"},
		{Name: "eks_capability", Service: "eks", Resource: "capability", Template: "arn:${Partition}:eks:${Region}:${Account}:capability/${ClusterName}/${CapabilityType}/${CapabilityName}/${UUID}"},
		{Name: "eks_cluster", Service: "eks", Resource: "cluster", Template: "arn:${Partition}:eks:${Region}:${Account}:cluster/${ClusterName}"},
		{Name: "eks_dashboard", Service: "eks", Resource: "dashboard", Template: "arn:${Partition}:eks:${Region}:${Account}:dashboard/${DashboardName}"},
		{Name: "eks_eks_anywhere_subscription", Service: "eks", Resource: "eks-anywhere-subscription", Template: "arn:${Partition}:eks:${Region}:${Account}:eks-anywhere-subscription/${UUID}"},
		{Name: "eks_fargateprofile", Service: "eks", Resource: "fargateprofile", Template: "arn:${Partition}:eks:${Region}:${Account}:fargateprofile/${ClusterName}/${FargateProfileName}/${UUID}"},
		{Name: "eks_identityproviderconfig", Service: "eks", Resource: "identityproviderconfig", Template: "arn:${Partition}:eks:${Region}:${Account}:identityproviderconfig/${ClusterName}/${IdentityProviderType}/${IdentityProviderConfigName}/${UUID}"},
		{Name: "eks_nodegroup", Service: "eks", Resource: "nodegroup", Template: "arn:${Partition}:eks:${Region}:${Account}:nodegroup/${ClusterName}/${NodegroupName}/${UUID}"},
		{Name: "eks_podidentityassociation", Service: "eks", Resource: "podidentityassociation", Template: "arn:${Partition}:eks:${Region}:${Account}:podidentityassociation/${ClusterName}/${UUID}"},
	})
}
