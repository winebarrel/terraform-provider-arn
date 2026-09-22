# arn:aws:account::111111111111:account/o-organization-id/member-account-id
output "account_account_in_organization" {
  value = provider::arn::account_account_in_organization("organization-id", "member-account-id")
}
