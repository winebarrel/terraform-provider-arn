// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: bcm-data-exports
// Source: https://servicereference.us-east-1.amazonaws.com/v1/bcm-data-exports/bcm-data-exports.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "bcm_data_exports_billingview", Service: "bcm-data-exports", Resource: "billingview", Template: "arn:${Partition}:billing::${Account}:billingview/${ResourceId}"},
		{Name: "bcm_data_exports_export", Service: "bcm-data-exports", Resource: "export", Template: "arn:${Partition}:bcm-data-exports:${Region}:${Account}:export/${Identifier}"},
		{Name: "bcm_data_exports_table", Service: "bcm-data-exports", Resource: "table", Template: "arn:${Partition}:bcm-data-exports:${Region}:${Account}:table/${Identifier}"},
	})
}
