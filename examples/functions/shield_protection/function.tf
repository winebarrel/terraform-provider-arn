# arn:aws:shield::111111111111:protection/id
output "shield_protection" {
  value = provider::arn::shield_protection("id")
}
