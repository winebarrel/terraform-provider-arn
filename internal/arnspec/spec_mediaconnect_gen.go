// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: mediaconnect
// Source: https://servicereference.us-east-1.amazonaws.com/v1/mediaconnect/mediaconnect.json
// Functions: 14
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "mediaconnect_bridge", Service: "mediaconnect", Resource: "Bridge", Template: "arn:${Partition}:mediaconnect:${Region}:${Account}:bridge:${BridgeId}:${BridgeName}"},
		{Name: "mediaconnect_entitlement", Service: "mediaconnect", Resource: "Entitlement", Template: "arn:${Partition}:mediaconnect:${Region}:${Account}:entitlement:${FlowId}:${EntitlementName}"},
		{Name: "mediaconnect_flow", Service: "mediaconnect", Resource: "Flow", Template: "arn:${Partition}:mediaconnect:${Region}:${Account}:flow:${FlowId}:${FlowName}"},
		{Name: "mediaconnect_gateway", Service: "mediaconnect", Resource: "Gateway", Template: "arn:${Partition}:mediaconnect:${Region}:${Account}:gateway:${GatewayId}:${GatewayName}"},
		{Name: "mediaconnect_gateway_instance", Service: "mediaconnect", Resource: "GatewayInstance", Template: "arn:${Partition}:mediaconnect:${Region}:${Account}:gateway:${GatewayId}:${GatewayName}:instance:${InstanceId}"},
		{Name: "mediaconnect_media_stream", Service: "mediaconnect", Resource: "MediaStream", Template: "arn:${Partition}:mediaconnect:${Region}:${Account}:flow:${FlowId}:${FlowName}/mediaStream/${MediaStreamName}"},
		{Name: "mediaconnect_offering", Service: "mediaconnect", Resource: "Offering", Template: "arn:${Partition}:mediaconnect:${Region}:offering:${OfferingId}"},
		{Name: "mediaconnect_output", Service: "mediaconnect", Resource: "Output", Template: "arn:${Partition}:mediaconnect:${Region}:${Account}:output:${OutputId}:${OutputName}"},
		{Name: "mediaconnect_reservation", Service: "mediaconnect", Resource: "Reservation", Template: "arn:${Partition}:mediaconnect:${Region}:${Account}:reservation:${ReservationId}:${ReservationName}"},
		{Name: "mediaconnect_router_input", Service: "mediaconnect", Resource: "RouterInput", Template: "arn:${Partition}:mediaconnect:${Region}:${Account}:routerInput:${RouterInputId}"},
		{Name: "mediaconnect_router_network_interface", Service: "mediaconnect", Resource: "RouterNetworkInterface", Template: "arn:${Partition}:mediaconnect:${Region}:${Account}:routerNetworkInterface:${RouterNetworkInterfaceId}"},
		{Name: "mediaconnect_router_output", Service: "mediaconnect", Resource: "RouterOutput", Template: "arn:${Partition}:mediaconnect:${Region}:${Account}:routerOutput:${RouterOutputId}"},
		{Name: "mediaconnect_source", Service: "mediaconnect", Resource: "Source", Template: "arn:${Partition}:mediaconnect:${Region}:${Account}:source:${SourceId}:${SourceName}"},
		{Name: "mediaconnect_vpc_interface", Service: "mediaconnect", Resource: "VpcInterface", Template: "arn:${Partition}:mediaconnect:${Region}:${Account}:flow:${FlowId}:${FlowName}/vpcInterface/${VpcInterfaceName}"},
	})
}
