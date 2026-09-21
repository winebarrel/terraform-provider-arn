// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: synthetics
// Source: https://servicereference.us-east-1.amazonaws.com/v1/synthetics/synthetics.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "synthetics_canary", Service: "synthetics", Resource: "canary", Template: "arn:${Partition}:synthetics:${Region}:${Account}:canary:${CanaryName}"},
		{Name: "synthetics_group", Service: "synthetics", Resource: "group", Template: "arn:${Partition}:synthetics:${Region}:${Account}:group:${GroupId}"},
	})
}
