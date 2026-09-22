# arn:aws:bedrock-agentcore:ap-northeast-1:aws:code-interpreter/code-interpreter-id
output "bedrock_agentcore_code_interpreter" {
  value = provider::arn::bedrock_agentcore_code_interpreter("code-interpreter-id")
}
