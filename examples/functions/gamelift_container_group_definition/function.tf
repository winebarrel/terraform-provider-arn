# arn:aws:gamelift:ap-northeast-1:111111111111:containergroupdefinition/name
output "gamelift_container_group_definition" {
  value = provider::arn::gamelift_container_group_definition("name")
}
