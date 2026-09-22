# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/metric/metric-id:metric-qualifier
output "connect_qualified_metric" {
  value = provider::arn::connect_qualified_metric("instance-id", "metric-id", "metric-qualifier")
}
