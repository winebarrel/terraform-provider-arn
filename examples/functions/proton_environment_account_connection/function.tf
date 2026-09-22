# arn:aws:proton:ap-northeast-1:111111111111:environment-account-connection/id
output "proton_environment_account_connection" {
  value = provider::arn::proton_environment_account_connection("id")
}
