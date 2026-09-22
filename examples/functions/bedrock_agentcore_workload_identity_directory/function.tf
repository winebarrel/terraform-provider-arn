# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:workload-identity-directory/directory-id
output "bedrock_agentcore_workload_identity_directory" {
  value = provider::arn::bedrock_agentcore_workload_identity_directory("directory-id")
}
