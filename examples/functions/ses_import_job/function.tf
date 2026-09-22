# arn:aws:ses:ap-northeast-1:111111111111:import-job/import-job-id
output "ses_import_job" {
  value = provider::arn::ses_import_job("import-job-id")
}
