# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/jobs/export/job-id
output "mobiletargeting_export_job" {
  value = provider::arn::mobiletargeting_export_job("app-id", "job-id")
}
