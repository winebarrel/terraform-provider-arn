# arn:aws:identitystore:::membership/*
output "identitystore_all_group_memberships" {
  value = provider::arn::identitystore_all_group_memberships()
}
