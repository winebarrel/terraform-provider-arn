# arn:aws:account-access:ap-northeast-1:111111111111:application/resource-id
output "account_access_application" {
  value = provider::arn::account_access_application("resource-id")
}
