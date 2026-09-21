// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: snow-device-management
// Source: https://servicereference.us-east-1.amazonaws.com/v1/snow-device-management/snow-device-management.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "snow_device_management_managed_device", Service: "snow-device-management", Resource: "managed-device", Template: "arn:${Partition}:snow-device-management:${Region}:${Account}:managed-device/${ResourceId}"},
		{Name: "snow_device_management_task", Service: "snow-device-management", Resource: "task", Template: "arn:${Partition}:snow-device-management:${Region}:${Account}:task/${ResourceId}"},
	})
}
