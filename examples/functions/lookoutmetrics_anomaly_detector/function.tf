# arn:aws:lookoutmetrics:ap-northeast-1:111111111111:AnomalyDetector:anomaly-detector-name
output "lookoutmetrics_anomaly_detector" {
  value = provider::arn::lookoutmetrics_anomaly_detector("anomaly-detector-name")
}
