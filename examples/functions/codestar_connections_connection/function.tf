# arn:aws:codestar-connections:ap-northeast-1:111111111111:connection/connection-id
output "codestar_connections_connection" {
  value = provider::arn::codestar_connections_connection("connection-id")
}
