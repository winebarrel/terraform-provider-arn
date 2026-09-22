# arn:aws:bedrock:ap-northeast-1:111111111111:model-evaluation-job/resource-id
output "bedrock_model_evaluation_job" {
  value = provider::arn::bedrock_model_evaluation_job("resource-id")
}
