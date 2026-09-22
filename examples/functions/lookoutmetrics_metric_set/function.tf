# arn:aws:lookoutmetrics:ap-northeast-1:111111111111:MetricSet/anomaly-detector-name/metric-set-name
output "lookoutmetrics_metric_set" {
  value = provider::arn::lookoutmetrics_metric_set("anomaly-detector-name", "metric-set-name")
}
