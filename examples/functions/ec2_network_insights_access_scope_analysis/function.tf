# arn:aws:ec2:ap-northeast-1:111111111111:network-insights-access-scope-analysis/network-insights-access-scope-analysis-id
output "ec2_network_insights_access_scope_analysis" {
  value = provider::arn::ec2_network_insights_access_scope_analysis("network-insights-access-scope-analysis-id")
}
