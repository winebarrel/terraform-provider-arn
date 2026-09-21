// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: payments
// Source: https://servicereference.us-east-1.amazonaws.com/v1/payments/payments.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "payments_payment_instrument", Service: "payments", Resource: "payment-instrument", Template: "arn:${Partition}:payments::${Account}:payment-instrument:${ResourceId}"},
	})
}
