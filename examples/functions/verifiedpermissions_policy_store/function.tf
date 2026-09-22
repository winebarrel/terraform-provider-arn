# arn:aws:verifiedpermissions::111111111111:policy-store/policy-store-id
output "verifiedpermissions_policy_store" {
  value = provider::arn::verifiedpermissions_policy_store("policy-store-id")
}
