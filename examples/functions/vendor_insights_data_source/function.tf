# arn:aws:vendor-insights:::data-source:resource-id
output "vendor_insights_data_source" {
  value = provider::arn::vendor_insights_data_source("resource-id")
}
