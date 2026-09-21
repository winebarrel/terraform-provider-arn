// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: eks-auth
// Source: https://servicereference.us-east-1.amazonaws.com/v1/eks-auth/eks-auth.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "eks_auth_cluster", Service: "eks-auth", Resource: "cluster", Template: "arn:${Partition}:eks:${Region}:${Account}:cluster/${ClusterName}"},
	})
}
