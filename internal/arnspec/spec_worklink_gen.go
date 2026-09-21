// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: worklink
// Source: https://servicereference.us-east-1.amazonaws.com/v1/worklink/worklink.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "worklink_fleet", Service: "worklink", Resource: "fleet", Template: "arn:${Partition}:worklink::${Account}:fleet/${FleetName}"},
	})
}
