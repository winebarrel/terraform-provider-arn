# arn:aws:ec2:ap-northeast-1:111111111111:network-insights-path/network-insights-path-id
output "ec2_network_insights_path" {
  value = provider::arn::ec2_network_insights_path("network-insights-path-id")
}
