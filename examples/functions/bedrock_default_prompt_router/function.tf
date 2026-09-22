# arn:aws:bedrock:ap-northeast-1:111111111111:default-prompt-router/resource-id
output "bedrock_default_prompt_router" {
  value = provider::arn::bedrock_default_prompt_router("resource-id")
}
