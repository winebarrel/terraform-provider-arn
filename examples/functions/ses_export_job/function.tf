# arn:aws:ses:ap-northeast-1:111111111111:export-job/export-job-id
output "ses_export_job" {
  value = provider::arn::ses_export_job("export-job-id")
}
