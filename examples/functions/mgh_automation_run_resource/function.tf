# arn:aws:mgh:ap-northeast-1:111111111111:automation-run/run-id
output "mgh_automation_run_resource" {
  value = provider::arn::mgh_automation_run_resource("run-id")
}
