# arn:aws:partnercentral::111111111111:catalog/catalog/ReportingData/table-id/Dashboard/dashboard-id
output "partnercentral_dashboard" {
  value = provider::arn::partnercentral_dashboard("catalog", "table-id", "dashboard-id")
}
