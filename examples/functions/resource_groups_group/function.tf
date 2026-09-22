# arn:aws:resource-groups:ap-northeast-1:111111111111:group/group-name
output "resource_groups_group" {
  value = provider::arn::resource_groups_group("group-name")
}
