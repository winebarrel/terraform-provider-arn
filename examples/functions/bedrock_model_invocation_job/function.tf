# arn:aws:bedrock:ap-northeast-1:111111111111:model-invocation-job/job-identifier
output "bedrock_model_invocation_job" {
  value = provider::arn::bedrock_model_invocation_job("job-identifier")
}
