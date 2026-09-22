# arn:aws:cloudhsm:ap-northeast-1:111111111111:backup/cloud-hsm-backup-instance-name
output "cloudhsm_backup" {
  value = provider::arn::cloudhsm_backup("cloud-hsm-backup-instance-name")
}
