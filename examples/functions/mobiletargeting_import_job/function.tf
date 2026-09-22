# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/jobs/import/job-id
output "mobiletargeting_import_job" {
  value = provider::arn::mobiletargeting_import_job("app-id", "job-id")
}
