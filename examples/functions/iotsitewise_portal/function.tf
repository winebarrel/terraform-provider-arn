# arn:aws:iotsitewise:ap-northeast-1:111111111111:portal/portal-id
output "iotsitewise_portal" {
  value = provider::arn::iotsitewise_portal("portal-id")
}
