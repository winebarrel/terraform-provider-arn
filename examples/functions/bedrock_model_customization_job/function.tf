# arn:aws:bedrock:ap-northeast-1:111111111111:model-customization-job/resource-id
output "bedrock_model_customization_job" {
  value = provider::arn::bedrock_model_customization_job("resource-id")
}
