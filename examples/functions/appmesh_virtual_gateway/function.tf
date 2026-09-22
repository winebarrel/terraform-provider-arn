# arn:aws:appmesh:ap-northeast-1:111111111111:mesh/mesh-name/virtualGateway/virtual-gateway-name
output "appmesh_virtual_gateway" {
  value = provider::arn::appmesh_virtual_gateway("mesh-name", "virtual-gateway-name")
}
