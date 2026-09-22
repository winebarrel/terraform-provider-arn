# arn:aws:bedrock:ap-northeast-1:111111111111:automated-reasoning-policy/automated-reasoning-policy-id:automated-reasoning-policy-version
output "bedrock_automated_reasoning_policy_version" {
  value = provider::arn::bedrock_automated_reasoning_policy_version("automated-reasoning-policy-id", "automated-reasoning-policy-version")
}
