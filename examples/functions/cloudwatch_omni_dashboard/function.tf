# arn:aws:cloudwatch:ap-northeast-1:111111111111:omni-dashboard/dashboard-id
output "cloudwatch_omni_dashboard" {
  value = provider::arn::cloudwatch_omni_dashboard("dashboard-id")
}
