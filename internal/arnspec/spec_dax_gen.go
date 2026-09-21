// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: dax
// Source: https://servicereference.us-east-1.amazonaws.com/v1/dax/dax.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "dax_application", Service: "dax", Resource: "application", Template: "arn:${Partition}:dax:${Region}:${Account}:cache/${ClusterName}"},
	})
}
