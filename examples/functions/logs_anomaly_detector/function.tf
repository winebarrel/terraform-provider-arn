# arn:aws:logs:ap-northeast-1:111111111111:anomaly-detector:detector-id
output "logs_anomaly_detector" {
  value = provider::arn::logs_anomaly_detector("detector-id")
}
