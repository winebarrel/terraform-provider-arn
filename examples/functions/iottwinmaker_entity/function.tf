# arn:aws:iottwinmaker:ap-northeast-1:111111111111:workspace/workspace-id/entity/entity-id
output "iottwinmaker_entity" {
  value = provider::arn::iottwinmaker_entity("workspace-id", "entity-id")
}
