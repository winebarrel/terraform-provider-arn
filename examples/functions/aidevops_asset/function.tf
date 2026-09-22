# arn:aws:aidevops:ap-northeast-1:111111111111:agentspace/agent-space-id/asset/asset-id
output "aidevops_asset" {
  value = provider::arn::aidevops_asset("agent-space-id", "asset-id")
}
