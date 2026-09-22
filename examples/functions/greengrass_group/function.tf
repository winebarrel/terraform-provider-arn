# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/groups/group-id
output "greengrass_group" {
  value = provider::arn::greengrass_group("group-id")
}
