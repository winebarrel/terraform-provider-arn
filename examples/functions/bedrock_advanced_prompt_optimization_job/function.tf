# arn:aws:bedrock:ap-northeast-1:111111111111:advanced-prompt-optimization-job/resource-id
output "bedrock_advanced_prompt_optimization_job" {
  value = provider::arn::bedrock_advanced_prompt_optimization_job("resource-id")
}
