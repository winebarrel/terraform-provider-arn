# arn:aws:identitystore::111111111111:identitystore/identity-store-id
output "identitystore_identitystore" {
  value = provider::arn::identitystore_identitystore("identity-store-id")
}
