# arn:aws:appmesh-preview:ap-northeast-1:111111111111:mesh/mesh-name
output "appmesh_preview_mesh" {
  value = provider::arn::appmesh_preview_mesh("mesh-name")
}
