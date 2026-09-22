# arn:aws:mediapackage:ap-northeast-1:111111111111:harvest_jobs/harvest-job-identifier
output "mediapackage_harvest_jobs" {
  value = provider::arn::mediapackage_harvest_jobs("harvest-job-identifier")
}
