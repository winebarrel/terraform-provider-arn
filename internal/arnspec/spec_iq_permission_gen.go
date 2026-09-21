// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: iq-permission
// Source: https://servicereference.us-east-1.amazonaws.com/v1/iq-permission/iq-permission.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "iq_permission_permission", Service: "iq-permission", Resource: "permission", Template: "arn:${Partition}:iq-permission:${Region}::permission/${PermissionRequestId}"},
	})
}
