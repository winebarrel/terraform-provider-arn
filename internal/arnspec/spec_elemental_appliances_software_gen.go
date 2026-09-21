// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: elemental-appliances-software
// Source: https://servicereference.us-east-1.amazonaws.com/v1/elemental-appliances-software/elemental-appliances-software.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "elemental_appliances_software_quote", Service: "elemental-appliances-software", Resource: "quote", Template: "arn:${Partition}:elemental-appliances-software:${Region}:${Account}:quote/${ResourceId}"},
	})
}
