# arn:aws:drs:ap-northeast-1:111111111111:recovery-plan-execution/recovery-plan-execution-id
output "drs_recovery_plan_execution_resource" {
  value = provider::arn::drs_recovery_plan_execution_resource("recovery-plan-execution-id")
}
