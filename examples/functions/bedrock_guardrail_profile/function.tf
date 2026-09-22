# arn:aws:bedrock:ap-northeast-1:111111111111:guardrail-profile/resource-id
output "bedrock_guardrail_profile" {
  value = provider::arn::bedrock_guardrail_profile("resource-id")
}
