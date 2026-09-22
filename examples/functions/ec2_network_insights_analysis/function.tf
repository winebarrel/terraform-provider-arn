# arn:aws:ec2:ap-northeast-1:111111111111:network-insights-analysis/network-insights-analysis-id
output "ec2_network_insights_analysis" {
  value = provider::arn::ec2_network_insights_analysis("network-insights-analysis-id")
}
