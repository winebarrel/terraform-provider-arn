# arn:aws:iotmanagedintegrations:ap-northeast-1:111111111111:account-association/account-association-id
output "iotmanagedintegrations_account_association" {
  value = provider::arn::iotmanagedintegrations_account_association("account-association-id")
}
