# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/groups/group-id/versions/version-id
output "greengrass_group_version" {
  value = provider::arn::greengrass_group_version("group-id", "version-id")
}
