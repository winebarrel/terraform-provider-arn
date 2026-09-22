# arn:aws:organizations::111111111111:account/o-organization-id/account-id
output "organizations_account" {
  value = provider::arn::organizations_account("organization-id", "account-id")
}
