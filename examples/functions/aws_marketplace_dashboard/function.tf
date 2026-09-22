# arn:aws:aws-marketplace::111111111111:catalog/ReportingData/fact-table/Dashboard/dashboard-name
output "aws_marketplace_dashboard" {
  value = provider::arn::aws_marketplace_dashboard("catalog", "fact-table", "dashboard-name")
}
