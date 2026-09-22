# arn:aws:bedrock:ap-northeast-1:111111111111:automated-reasoning-policy/automated-reasoning-policy-id
output "bedrock_automated_reasoning_policy" {
  value = provider::arn::bedrock_automated_reasoning_policy("automated-reasoning-policy-id")
}
