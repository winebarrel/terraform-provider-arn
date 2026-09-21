// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: billing
// Source: https://servicereference.us-east-1.amazonaws.com/v1/billing/billing.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "billing_billingview", Service: "billing", Resource: "billingview", Template: "arn:${Partition}:billing::${Account}:billingview/${ResourceId}"},
	})
}
