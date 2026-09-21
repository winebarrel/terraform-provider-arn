// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: appmesh-preview
// Source: https://servicereference.us-east-1.amazonaws.com/v1/appmesh-preview/appmesh-preview.json
// Functions: 7
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "appmesh_preview_gateway_route", Service: "appmesh-preview", Resource: "gatewayRoute", Template: "arn:${Partition}:appmesh-preview:${Region}:${Account}:mesh/${MeshName}/virtualGateway/${VirtualGatewayName}/gatewayRoute/${GatewayRouteName}"},
		{Name: "appmesh_preview_mesh", Service: "appmesh-preview", Resource: "mesh", Template: "arn:${Partition}:appmesh-preview:${Region}:${Account}:mesh/${MeshName}"},
		{Name: "appmesh_preview_route", Service: "appmesh-preview", Resource: "route", Template: "arn:${Partition}:appmesh-preview:${Region}:${Account}:mesh/${MeshName}/virtualRouter/${VirtualRouterName}/route/${RouteName}"},
		{Name: "appmesh_preview_virtual_gateway", Service: "appmesh-preview", Resource: "virtualGateway", Template: "arn:${Partition}:appmesh-preview:${Region}:${Account}:mesh/${MeshName}/virtualGateway/${VirtualGatewayName}"},
		{Name: "appmesh_preview_virtual_node", Service: "appmesh-preview", Resource: "virtualNode", Template: "arn:${Partition}:appmesh-preview:${Region}:${Account}:mesh/${MeshName}/virtualNode/${VirtualNodeName}"},
		{Name: "appmesh_preview_virtual_router", Service: "appmesh-preview", Resource: "virtualRouter", Template: "arn:${Partition}:appmesh-preview:${Region}:${Account}:mesh/${MeshName}/virtualRouter/${VirtualRouterName}"},
		{Name: "appmesh_preview_virtual_service", Service: "appmesh-preview", Resource: "virtualService", Template: "arn:${Partition}:appmesh-preview:${Region}:${Account}:mesh/${MeshName}/virtualService/${VirtualServiceName}"},
	})
}
