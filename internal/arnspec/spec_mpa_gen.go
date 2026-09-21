// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: mpa
// Source: https://servicereference.us-east-1.amazonaws.com/v1/mpa/mpa.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "mpa_approval_team", Service: "mpa", Resource: "approval-team", Template: "arn:${Partition}:mpa:${Region}:${Account}:approval-team/${ApprovalTeamId}"},
		{Name: "mpa_identity_source", Service: "mpa", Resource: "identity-source", Template: "arn:${Partition}:mpa:${Region}:${Account}:identity-source/${IdentitySourceId}"},
		{Name: "mpa_session", Service: "mpa", Resource: "session", Template: "arn:${Partition}:mpa:${Region}:${Account}:session/${SessionId}"},
	})
}
