// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: servicequotas
// Source: https://servicereference.us-east-1.amazonaws.com/v1/servicequotas/servicequotas.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "servicequotas_quota", Service: "servicequotas", Resource: "quota", Template: "arn:${Partition}:servicequotas:${Region}:${Account}:${ServiceCode}/${QuotaCode}"},
	})
}
