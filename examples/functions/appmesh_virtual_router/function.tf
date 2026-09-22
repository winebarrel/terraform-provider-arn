# arn:aws:appmesh:ap-northeast-1:111111111111:mesh/mesh-name/virtualRouter/virtual-router-name
output "appmesh_virtual_router" {
  value = provider::arn::appmesh_virtual_router("mesh-name", "virtual-router-name")
}
