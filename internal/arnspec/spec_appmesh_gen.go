// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: appmesh
// Source: https://servicereference.us-east-1.amazonaws.com/v1/appmesh/appmesh.json
// Functions: 7
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "appmesh_gateway_route", Service: "appmesh", Resource: "gatewayRoute", Template: "arn:${Partition}:appmesh:${Region}:${Account}:mesh/${MeshName}/virtualGateway/${VirtualGatewayName}/gatewayRoute/${GatewayRouteName}"},
		{Name: "appmesh_mesh", Service: "appmesh", Resource: "mesh", Template: "arn:${Partition}:appmesh:${Region}:${Account}:mesh/${MeshName}"},
		{Name: "appmesh_route", Service: "appmesh", Resource: "route", Template: "arn:${Partition}:appmesh:${Region}:${Account}:mesh/${MeshName}/virtualRouter/${VirtualRouterName}/route/${RouteName}"},
		{Name: "appmesh_virtual_gateway", Service: "appmesh", Resource: "virtualGateway", Template: "arn:${Partition}:appmesh:${Region}:${Account}:mesh/${MeshName}/virtualGateway/${VirtualGatewayName}"},
		{Name: "appmesh_virtual_node", Service: "appmesh", Resource: "virtualNode", Template: "arn:${Partition}:appmesh:${Region}:${Account}:mesh/${MeshName}/virtualNode/${VirtualNodeName}"},
		{Name: "appmesh_virtual_router", Service: "appmesh", Resource: "virtualRouter", Template: "arn:${Partition}:appmesh:${Region}:${Account}:mesh/${MeshName}/virtualRouter/${VirtualRouterName}"},
		{Name: "appmesh_virtual_service", Service: "appmesh", Resource: "virtualService", Template: "arn:${Partition}:appmesh:${Region}:${Account}:mesh/${MeshName}/virtualService/${VirtualServiceName}"},
	})
}
