# arn:aws:verifiedpermissions:ap-northeast-1:111111111111:policy-store-alias/alias-name
output "verifiedpermissions_policy_store_alias" {
  value = provider::arn::verifiedpermissions_policy_store_alias("alias-name")
}
