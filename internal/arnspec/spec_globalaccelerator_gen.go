// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: globalaccelerator
// Source: https://servicereference.us-east-1.amazonaws.com/v1/globalaccelerator/globalaccelerator.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "globalaccelerator_accelerator", Service: "globalaccelerator", Resource: "accelerator", Template: "arn:${Partition}:globalaccelerator::${Account}:accelerator/${ResourceId}"},
		{Name: "globalaccelerator_attachment", Service: "globalaccelerator", Resource: "attachment", Template: "arn:${Partition}:globalaccelerator::${Account}:attachment/${ResourceId}"},
		{Name: "globalaccelerator_endpointgroup", Service: "globalaccelerator", Resource: "endpointgroup", Template: "arn:${Partition}:globalaccelerator::${Account}:accelerator/${ResourceId}/listener/${ListenerId}/endpoint-group/${EndpointGroupId}"},
		{Name: "globalaccelerator_listener", Service: "globalaccelerator", Resource: "listener", Template: "arn:${Partition}:globalaccelerator::${Account}:accelerator/${ResourceId}/listener/${ListenerId}"},
	})
}
