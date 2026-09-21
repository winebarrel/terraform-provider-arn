// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: rolesanywhere
// Source: https://servicereference.us-east-1.amazonaws.com/v1/rolesanywhere/rolesanywhere.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "rolesanywhere_crl", Service: "rolesanywhere", Resource: "crl", Template: "arn:${Partition}:rolesanywhere:${Region}:${Account}:crl/${CrlId}"},
		{Name: "rolesanywhere_profile", Service: "rolesanywhere", Resource: "profile", Template: "arn:${Partition}:rolesanywhere:${Region}:${Account}:profile/${ProfileId}"},
		{Name: "rolesanywhere_subject", Service: "rolesanywhere", Resource: "subject", Template: "arn:${Partition}:rolesanywhere:${Region}:${Account}:subject/${SubjectId}"},
		{Name: "rolesanywhere_trust_anchor", Service: "rolesanywhere", Resource: "trust-anchor", Template: "arn:${Partition}:rolesanywhere:${Region}:${Account}:trust-anchor/${TrustAnchorId}"},
	})
}
