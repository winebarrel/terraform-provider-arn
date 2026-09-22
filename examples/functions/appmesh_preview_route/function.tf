# arn:aws:appmesh-preview:ap-northeast-1:111111111111:mesh/mesh-name/virtualRouter/virtual-router-name/route/route-name
output "appmesh_preview_route" {
  value = provider::arn::appmesh_preview_route("mesh-name", "virtual-router-name", "route-name")
}
