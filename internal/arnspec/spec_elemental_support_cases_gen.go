// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: elemental-support-cases
// Source: https://servicereference.us-east-1.amazonaws.com/v1/elemental-support-cases/elemental-support-cases.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "elemental_support_cases_case", Service: "elemental-support-cases", Resource: "case", Template: "arn:${Partition}:elemental-support-cases::${Account}:case/${ResourceId}"},
	})
}
