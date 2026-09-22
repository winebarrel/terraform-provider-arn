# arn:aws:bedrock:ap-northeast-1:111111111111:prompt-router/resource-id
output "bedrock_prompt_router" {
  value = provider::arn::bedrock_prompt_router("resource-id")
}
