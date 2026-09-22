# arn:aws:iot:ap-northeast-1:111111111111:job/job-id
output "iot_job" {
  value = provider::arn::iot_job("job-id")
}
