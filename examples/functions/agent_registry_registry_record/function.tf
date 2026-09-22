# arn:aws:agent-registry:ap-northeast-1:111111111111:registry/registry-id/record/record-id
output "agent_registry_registry_record" {
  value = provider::arn::agent_registry_registry_record("registry-id", "record-id")
}
