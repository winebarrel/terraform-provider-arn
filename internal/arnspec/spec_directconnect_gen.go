// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: directconnect
// Source: https://servicereference.us-east-1.amazonaws.com/v1/directconnect/directconnect.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "directconnect_dx_gateway", Service: "directconnect", Resource: "dx-gateway", Template: "arn:${Partition}:directconnect::${Account}:dx-gateway/${DirectConnectGatewayId}"},
		{Name: "directconnect_dx_resiliency_group", Service: "directconnect", Resource: "dx-resiliency-group", Template: "arn:${Partition}:directconnect::${Account}:dx-resiliency-group/${ResiliencyGroupId}"},
		{Name: "directconnect_dxcon", Service: "directconnect", Resource: "dxcon", Template: "arn:${Partition}:directconnect:${Region}:${Account}:dxcon/${ConnectionId}"},
		{Name: "directconnect_dxlag", Service: "directconnect", Resource: "dxlag", Template: "arn:${Partition}:directconnect:${Region}:${Account}:dxlag/${LagId}"},
		{Name: "directconnect_dxvif", Service: "directconnect", Resource: "dxvif", Template: "arn:${Partition}:directconnect:${Region}:${Account}:dxvif/${VirtualInterfaceId}"},
	})
}
