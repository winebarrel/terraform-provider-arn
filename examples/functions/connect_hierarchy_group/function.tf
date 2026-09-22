# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/agent-group/hierarchy-group-id
output "connect_hierarchy_group" {
  value = provider::arn::connect_hierarchy_group("instance-id", "hierarchy-group-id")
}
