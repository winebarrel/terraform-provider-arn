# arn:aws:shield::111111111111:protection-group/id
output "shield_protection_group" {
  value = provider::arn::shield_protection_group("id")
}
