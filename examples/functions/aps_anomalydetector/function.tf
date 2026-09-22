# arn:aws:aps:ap-northeast-1:111111111111:anomalydetector/workspace-id/anomaly-detector-id
output "aps_anomalydetector" {
  value = provider::arn::aps_anomalydetector("workspace-id", "anomaly-detector-id")
}
