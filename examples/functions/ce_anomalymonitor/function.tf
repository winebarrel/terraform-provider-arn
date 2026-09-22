# arn:aws:ce::111111111111:anomalymonitor/identifier
output "ce_anomalymonitor" {
  value = provider::arn::ce_anomalymonitor("identifier")
}
