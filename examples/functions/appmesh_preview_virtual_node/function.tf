# arn:aws:appmesh-preview:ap-northeast-1:111111111111:mesh/mesh-name/virtualNode/virtual-node-name
output "appmesh_preview_virtual_node" {
  value = provider::arn::appmesh_preview_virtual_node("mesh-name", "virtual-node-name")
}
