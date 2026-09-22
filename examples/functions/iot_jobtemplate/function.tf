# arn:aws:iot:ap-northeast-1:111111111111:jobtemplate/job-template-id
output "iot_jobtemplate" {
  value = provider::arn::iot_jobtemplate("job-template-id")
}
