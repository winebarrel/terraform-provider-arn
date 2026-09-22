# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:batch-evaluate/batch-evaluation-id
output "bedrock_agentcore_batch_evaluate" {
  value = provider::arn::bedrock_agentcore_batch_evaluate("batch-evaluation-id")
}
