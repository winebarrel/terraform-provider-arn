# arn:aws:cloudwatch:ap-northeast-1:111111111111:metric-stream/metric-stream-name
output "cloudwatch_metric_stream" {
  value = provider::arn::cloudwatch_metric_stream("metric-stream-name")
}
