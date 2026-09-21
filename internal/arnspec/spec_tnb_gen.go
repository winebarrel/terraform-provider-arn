// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: tnb
// Source: https://servicereference.us-east-1.amazonaws.com/v1/tnb/tnb.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "tnb_function_instance", Service: "tnb", Resource: "function-instance", Template: "arn:${Partition}:tnb:${Region}:${Account}:function-instance/${FunctionInstanceId}"},
		{Name: "tnb_function_package", Service: "tnb", Resource: "function-package", Template: "arn:${Partition}:tnb:${Region}:${Account}:function-package/${FunctionPackageId}"},
		{Name: "tnb_network_instance", Service: "tnb", Resource: "network-instance", Template: "arn:${Partition}:tnb:${Region}:${Account}:network-instance/${NetworkInstanceId}"},
		{Name: "tnb_network_operation", Service: "tnb", Resource: "network-operation", Template: "arn:${Partition}:tnb:${Region}:${Account}:network-operation/${NetworkOperationId}"},
		{Name: "tnb_network_package", Service: "tnb", Resource: "network-package", Template: "arn:${Partition}:tnb:${Region}:${Account}:network-package/${NetworkPackageId}"},
	})
}
