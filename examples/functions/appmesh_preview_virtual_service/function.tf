# arn:aws:appmesh-preview:ap-northeast-1:111111111111:mesh/mesh-name/virtualService/virtual-service-name
output "appmesh_preview_virtual_service" {
  value = provider::arn::appmesh_preview_virtual_service("mesh-name", "virtual-service-name")
}
