# arn:aws:identitystore:::group/group-id
output "identitystore_group" {
  value = provider::arn::identitystore_group("group-id")
}
