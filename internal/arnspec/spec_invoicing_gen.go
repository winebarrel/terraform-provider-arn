// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: invoicing
// Source: https://servicereference.us-east-1.amazonaws.com/v1/invoicing/invoicing.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "invoicing_invoice_unit", Service: "invoicing", Resource: "invoice-unit", Template: "arn:${Partition}:invoicing::${Account}:invoice-unit/${Identifier}"},
		{Name: "invoicing_procurement_portal_preference", Service: "invoicing", Resource: "procurement-portal-preference", Template: "arn:${Partition}:invoicing::${Account}:procurement-portal-preference/${Identifier}"},
	})
}
