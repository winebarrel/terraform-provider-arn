# arn:aws:appflow:ap-northeast-1:111111111111:connector/connector-label
output "appflow_connector" {
  value = provider::arn::appflow_connector("connector-label")
}
