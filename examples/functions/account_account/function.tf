# arn:aws:account::111111111111:account
output "account_account" {
  value = provider::arn::account_account()
}
