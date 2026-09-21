// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: arc-region-switch
// Source: https://servicereference.us-east-1.amazonaws.com/v1/arc-region-switch/arc-region-switch.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "arc_region_switch_plan", Service: "arc-region-switch", Resource: "plan", Template: "arn:${Partition}:arc-region-switch::${Account}:plan/${ResourceId}"},
	})
}
