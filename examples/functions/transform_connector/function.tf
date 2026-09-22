# arn:aws:transform:ap-northeast-1:111111111111:connector/workspace-id/connector-id
output "transform_connector" {
  value = provider::arn::transform_connector("workspace-id", "connector-id")
}
