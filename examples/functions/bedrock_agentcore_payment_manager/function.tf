# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:payment-manager/payment-manager-id
output "bedrock_agentcore_payment_manager" {
  value = provider::arn::bedrock_agentcore_payment_manager("payment-manager-id")
}
