# arn:aws:cloudtrail:ap-northeast-1:111111111111:dashboard/dashboard-name
output "cloudtrail_dashboard" {
  value = provider::arn::cloudtrail_dashboard("dashboard-name")
}
