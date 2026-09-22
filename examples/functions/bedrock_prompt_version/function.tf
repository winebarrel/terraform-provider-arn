# arn:aws:bedrock:ap-northeast-1:111111111111:prompt/prompt-id:prompt-version
output "bedrock_prompt_version" {
  value = provider::arn::bedrock_prompt_version("prompt-id", "prompt-version")
}
