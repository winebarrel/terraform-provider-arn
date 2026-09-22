# arn:aws:identitystore:::user/user-id
output "identitystore_user" {
  value = provider::arn::identitystore_user("user-id")
}
