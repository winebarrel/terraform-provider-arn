// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: rbin
// Source: https://servicereference.us-east-1.amazonaws.com/v1/rbin/rbin.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "rbin_rule", Service: "rbin", Resource: "rule", Template: "arn:${Partition}:rbin:${Region}:${Account}:rule/${ResourceName}"},
	})
}
