# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/metric/metric-id
output "connect_metric" {
  value = provider::arn::connect_metric("instance-id", "metric-id")
}
