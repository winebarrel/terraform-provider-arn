# arn:aws:apprunner:ap-northeast-1:111111111111:connection/connection-name/connection-id
output "apprunner_connection" {
  value = provider::arn::apprunner_connection("connection-name", "connection-id")
}
