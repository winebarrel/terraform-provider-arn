# arn:aws:aws-marketplace::111111111111:catalog/ReportingData/fact-table/Dashboard/dashboard-name
output "aws_marketplace_seller_dashboard" {
  value = provider::arn::aws_marketplace_seller_dashboard("catalog", "fact-table", "dashboard-name")
}
