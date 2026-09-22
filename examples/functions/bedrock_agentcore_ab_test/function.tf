# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:ab-test/ab-test-id
output "bedrock_agentcore_ab_test" {
  value = provider::arn::bedrock_agentcore_ab_test("ab-test-id")
}
