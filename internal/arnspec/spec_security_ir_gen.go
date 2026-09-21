// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: security-ir
// Source: https://servicereference.us-east-1.amazonaws.com/v1/security-ir/security-ir.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "security_ir_case", Service: "security-ir", Resource: "case", Template: "arn:${Partition}:security-ir:${Region}:${Account}:case/${CaseId}"},
		{Name: "security_ir_membership", Service: "security-ir", Resource: "membership", Template: "arn:${Partition}:security-ir:${Region}:${Account}:membership/${MembershipId}"},
	})
}
