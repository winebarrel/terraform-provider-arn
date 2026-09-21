// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: trustedadvisor
// Source: https://servicereference.us-east-1.amazonaws.com/v1/trustedadvisor/trustedadvisor.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "trustedadvisor_checks", Service: "trustedadvisor", Resource: "checks", Template: "arn:${Partition}:trustedadvisor:${Region}:${Account}:checks/${CategoryCode}/${CheckId}"},
	})
}
