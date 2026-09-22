# arn:aws:cloudwatch::111111111111:dashboard/dashboard-name
output "cloudwatch_dashboard" {
  value = provider::arn::cloudwatch_dashboard("dashboard-name")
}
