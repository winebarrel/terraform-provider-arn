# arn:aws:forecast:::algorithm/resource-id
output "forecast_algorithm" {
  value = provider::arn::forecast_algorithm("resource-id")
}
