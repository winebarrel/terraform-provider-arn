# arn:aws:backup-gateway:ap-northeast-1:111111111111:hypervisor/hypervisor-id
output "backup_gateway_hypervisor" {
  value = provider::arn::backup_gateway_hypervisor("hypervisor-id")
}
