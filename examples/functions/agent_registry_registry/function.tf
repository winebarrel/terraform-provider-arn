# arn:aws:agent-registry:ap-northeast-1:111111111111:registry/registry-id
output "agent_registry_registry" {
  value = provider::arn::agent_registry_registry("registry-id")
}
