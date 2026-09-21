// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: xray
// Source: https://servicereference.us-east-1.amazonaws.com/v1/xray/xray.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "xray_group", Service: "xray", Resource: "group", Template: "arn:${Partition}:xray:${Region}:${Account}:group/${GroupName}/${Id}"},
		{Name: "xray_sampling_rule", Service: "xray", Resource: "sampling-rule", Template: "arn:${Partition}:xray:${Region}:${Account}:sampling-rule/${SamplingRuleName}"},
	})
}
