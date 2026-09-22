# arn:aws:iottwinmaker:ap-northeast-1:111111111111:workspace/workspace-id/component-type/component-type-id
output "iottwinmaker_component_type" {
  value = provider::arn::iottwinmaker_component_type("workspace-id", "component-type-id")
}
