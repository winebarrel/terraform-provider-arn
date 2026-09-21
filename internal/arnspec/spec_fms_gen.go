// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: fms
// Source: https://servicereference.us-east-1.amazonaws.com/v1/fms/fms.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "fms_applications_list", Service: "fms", Resource: "applications-list", Template: "arn:${Partition}:fms:${Region}:${Account}:applications-list/${Id}"},
		{Name: "fms_policy", Service: "fms", Resource: "policy", Template: "arn:${Partition}:fms:${Region}:${Account}:policy/${Id}"},
		{Name: "fms_protocols_list", Service: "fms", Resource: "protocols-list", Template: "arn:${Partition}:fms:${Region}:${Account}:protocols-list/${Id}"},
		{Name: "fms_resource_set", Service: "fms", Resource: "resource-set", Template: "arn:${Partition}:fms:${Region}:${Account}:resource-set/${Id}"},
	})
}
