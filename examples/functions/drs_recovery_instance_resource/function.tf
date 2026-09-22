# arn:aws:drs:ap-northeast-1:111111111111:recovery-instance/recovery-instance-id
output "drs_recovery_instance_resource" {
  value = provider::arn::drs_recovery_instance_resource("recovery-instance-id")
}
