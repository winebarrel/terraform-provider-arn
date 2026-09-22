# arn:aws:ce::111111111111:anomalysubscription/identifier
output "ce_anomalysubscription" {
  value = provider::arn::ce_anomalysubscription("identifier")
}
