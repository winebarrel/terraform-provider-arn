# arn:aws:codestar-connections:ap-northeast-1:111111111111:host/host-id
output "codestar_connections_host" {
  value = provider::arn::codestar_connections_host("host-id")
}
