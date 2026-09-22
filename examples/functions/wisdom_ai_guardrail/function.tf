# arn:aws:wisdom:ap-northeast-1:111111111111:ai-guardrail/assistant-id/ai-guardrail-id
output "wisdom_ai_guardrail" {
  value = provider::arn::wisdom_ai_guardrail("assistant-id", "ai-guardrail-id")
}
