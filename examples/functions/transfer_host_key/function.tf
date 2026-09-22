# arn:aws:transfer:ap-northeast-1:111111111111:host-key/server-id/host-key-id
output "transfer_host_key" {
  value = provider::arn::transfer_host_key("server-id", "host-key-id")
}
