# arn:aws:geo:ap-northeast-1:111111111111:job/job-id
output "geo_job" {
  value = provider::arn::geo_job("job-id")
}
