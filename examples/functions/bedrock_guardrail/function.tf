# arn:aws:bedrock:ap-northeast-1:111111111111:guardrail/guardrail-id
output "bedrock_guardrail" {
  value = provider::arn::bedrock_guardrail("guardrail-id")
}
