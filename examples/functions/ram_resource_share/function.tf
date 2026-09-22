# arn:aws:ram:ap-northeast-1:111111111111:resource-share/resource-path
output "ram_resource_share" {
  value = provider::arn::ram_resource_share("resource-path")
}
