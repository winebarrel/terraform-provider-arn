# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:code-interpreter-custom/code-interpreter-id
output "bedrock_agentcore_code_interpreter_custom" {
  value = provider::arn::bedrock_agentcore_code_interpreter_custom("code-interpreter-id")
}
