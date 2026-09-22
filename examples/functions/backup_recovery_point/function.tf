# arn:aws:vendor:ap-northeast-1:*:resource-type:recovery-point-id
output "backup_recovery_point" {
  value = provider::arn::backup_recovery_point("vendor", "resource-type", "recovery-point-id")
}
