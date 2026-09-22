# arn:aws:appstudio:ap-northeast-1:111111111111:connector/connection-id
output "appstudio_connector" {
  value = provider::arn::appstudio_connector("connection-id")
}
