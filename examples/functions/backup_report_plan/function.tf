# arn:aws:backup:ap-northeast-1:111111111111:report-plan:report-plan-name-report-plan-id
output "backup_report_plan" {
  value = provider::arn::backup_report_plan("report-plan-name", "report-plan-id")
}
