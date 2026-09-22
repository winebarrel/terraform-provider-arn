# arn:aws:bedrock:ap-northeast-1:111111111111:flow/flow-id/alias/flow-alias-id
output "bedrock_flow_alias" {
  value = provider::arn::bedrock_flow_alias("flow-id", "flow-alias-id")
}
