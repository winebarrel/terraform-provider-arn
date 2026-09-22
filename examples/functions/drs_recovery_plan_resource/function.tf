# arn:aws:drs:ap-northeast-1:111111111111:recovery-plan/recovery-plan-id
output "drs_recovery_plan_resource" {
  value = provider::arn::drs_recovery_plan_resource("recovery-plan-id")
}
