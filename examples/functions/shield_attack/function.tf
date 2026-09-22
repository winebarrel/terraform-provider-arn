# arn:aws:shield::111111111111:attack/id
output "shield_attack" {
  value = provider::arn::shield_attack("id")
}
