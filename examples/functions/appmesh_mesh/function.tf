# arn:aws:appmesh:ap-northeast-1:111111111111:mesh/mesh-name
output "appmesh_mesh" {
  value = provider::arn::appmesh_mesh("mesh-name")
}
