# arn:aws:securityhub:ap-northeast-1:111111111111:aggregatorv2/aggregator-v2-id
output "securityhub_aggregatorv2" {
  value = provider::arn::securityhub_aggregatorv2("aggregator-v2-id")
}
