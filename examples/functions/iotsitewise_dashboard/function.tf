# arn:aws:iotsitewise:ap-northeast-1:111111111111:dashboard/dashboard-id
output "iotsitewise_dashboard" {
  value = provider::arn::iotsitewise_dashboard("dashboard-id")
}
