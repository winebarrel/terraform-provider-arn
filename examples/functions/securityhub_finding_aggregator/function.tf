# arn:aws:securityhub:ap-northeast-1:111111111111:finding-aggregator/finding-aggregator-id
output "securityhub_finding_aggregator" {
  value = provider::arn::securityhub_finding_aggregator("finding-aggregator-id")
}
