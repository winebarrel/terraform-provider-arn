# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:workload-identity-directory/directory-id/workload-identity/workload-identity-name
output "bedrock_agentcore_workload_identity" {
  value = provider::arn::bedrock_agentcore_workload_identity("directory-id", "workload-identity-name")
}
