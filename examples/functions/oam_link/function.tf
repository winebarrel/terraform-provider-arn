# arn:aws:oam:ap-northeast-1:111111111111:link/resource-id
output "oam_link" {
  value = provider::arn::oam_link("resource-id")
}
