# arn:aws:backup-gateway:ap-northeast-1:111111111111:vm/virtualmachine-id
output "backup_gateway_virtualmachine" {
  value = provider::arn::backup_gateway_virtualmachine("virtualmachine-id")
}
