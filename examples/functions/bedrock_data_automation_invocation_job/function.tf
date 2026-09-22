# arn:aws:bedrock:ap-northeast-1:111111111111:data-automation-invocation/job-id
output "bedrock_data_automation_invocation_job" {
  value = provider::arn::bedrock_data_automation_invocation_job("job-id")
}
