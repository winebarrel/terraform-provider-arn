// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: managedblockchain
// Source: https://servicereference.us-east-1.amazonaws.com/v1/managedblockchain/managedblockchain.json
// Functions: 6
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "managedblockchain_accessor", Service: "managedblockchain", Resource: "accessor", Template: "arn:${Partition}:managedblockchain:${Region}:${Account}:accessors/${AccessorId}"},
		{Name: "managedblockchain_invitation", Service: "managedblockchain", Resource: "invitation", Template: "arn:${Partition}:managedblockchain:${Region}:${Account}:invitations/${InvitationId}"},
		{Name: "managedblockchain_member", Service: "managedblockchain", Resource: "member", Template: "arn:${Partition}:managedblockchain:${Region}:${Account}:members/${MemberId}"},
		{Name: "managedblockchain_network", Service: "managedblockchain", Resource: "network", Template: "arn:${Partition}:managedblockchain:${Region}::networks/${NetworkId}"},
		{Name: "managedblockchain_node", Service: "managedblockchain", Resource: "node", Template: "arn:${Partition}:managedblockchain:${Region}:${Account}:nodes/${NodeId}"},
		{Name: "managedblockchain_proposal", Service: "managedblockchain", Resource: "proposal", Template: "arn:${Partition}:managedblockchain:${Region}::proposals/${ProposalId}"},
	})
}
