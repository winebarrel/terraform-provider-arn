# arn:aws:bedrock:ap-northeast-1:111111111111:model-copy-job/resource-id
output "bedrock_model_copy_job" {
  value = provider::arn::bedrock_model_copy_job("resource-id")
}
