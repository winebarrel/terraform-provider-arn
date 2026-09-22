# arn:aws:appmesh:ap-northeast-1:111111111111:mesh/mesh-name/virtualGateway/virtual-gateway-name/gatewayRoute/gateway-route-name
output "appmesh_gateway_route" {
  value = provider::arn::appmesh_gateway_route("mesh-name", "virtual-gateway-name", "gateway-route-name")
}
