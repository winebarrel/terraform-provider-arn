# arn:aws:sso:::account/account-id
output "sso_account" {
  value = provider::arn::sso_account("account-id")
}
