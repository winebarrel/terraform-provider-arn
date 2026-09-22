# arn:aws:bedrock:ap-northeast-1:111111111111:prompt/prompt-id
output "bedrock_prompt" {
  value = provider::arn::bedrock_prompt("prompt-id")
}
