# arn:aws:transfer:ap-northeast-1:111111111111:server/server-id
output "transfer_server" {
  value = provider::arn::transfer_server("server-id")
}
