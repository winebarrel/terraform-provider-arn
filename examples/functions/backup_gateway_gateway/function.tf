# arn:aws:backup-gateway:ap-northeast-1:111111111111:gateway/gateway-id
output "backup_gateway_gateway" {
  value = provider::arn::backup_gateway_gateway("gateway-id")
}
