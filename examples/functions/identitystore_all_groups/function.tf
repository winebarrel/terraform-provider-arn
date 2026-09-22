# arn:aws:identitystore:::group/*
output "identitystore_all_groups" {
  value = provider::arn::identitystore_all_groups()
}
