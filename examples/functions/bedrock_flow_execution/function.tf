# arn:aws:bedrock:ap-northeast-1:111111111111:flow/flow-id/alias/flow-alias-id/execution/flow-execution-id
output "bedrock_flow_execution" {
  value = provider::arn::bedrock_flow_execution("flow-id", "flow-alias-id", "flow-execution-id")
}
