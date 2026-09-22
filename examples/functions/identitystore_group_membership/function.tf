# arn:aws:identitystore:::membership/membership-id
output "identitystore_group_membership" {
  value = provider::arn::identitystore_group_membership("membership-id")
}
