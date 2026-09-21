// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: ce
// Source: https://servicereference.us-east-1.amazonaws.com/v1/ce/ce.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "ce_anomalymonitor", Service: "ce", Resource: "anomalymonitor", Template: "arn:${Partition}:ce::${Account}:anomalymonitor/${Identifier}"},
		{Name: "ce_anomalysubscription", Service: "ce", Resource: "anomalysubscription", Template: "arn:${Partition}:ce::${Account}:anomalysubscription/${Identifier}"},
		{Name: "ce_billingview", Service: "ce", Resource: "billingview", Template: "arn:${Partition}:billing::${Account}:billingview/${ResourceId}"},
		{Name: "ce_costcategory", Service: "ce", Resource: "costcategory", Template: "arn:${Partition}:ce::${Account}:costcategory/${Identifier}"},
	})
}
