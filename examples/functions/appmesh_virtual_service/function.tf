# arn:aws:appmesh:ap-northeast-1:111111111111:mesh/mesh-name/virtualService/virtual-service-name
output "appmesh_virtual_service" {
  value = provider::arn::appmesh_virtual_service("mesh-name", "virtual-service-name")
}
