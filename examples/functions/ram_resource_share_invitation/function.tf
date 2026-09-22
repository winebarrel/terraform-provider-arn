# arn:aws:ram:ap-northeast-1:111111111111:resource-share-invitation/resource-path
output "ram_resource_share_invitation" {
  value = provider::arn::ram_resource_share_invitation("resource-path")
}
