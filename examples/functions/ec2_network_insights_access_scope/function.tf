# arn:aws:ec2:ap-northeast-1:111111111111:network-insights-access-scope/network-insights-access-scope-id
output "ec2_network_insights_access_scope" {
  value = provider::arn::ec2_network_insights_access_scope("network-insights-access-scope-id")
}
