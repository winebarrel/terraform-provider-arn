# arn:aws:identitystore:::user/*
output "identitystore_all_users" {
  value = provider::arn::identitystore_all_users()
}
