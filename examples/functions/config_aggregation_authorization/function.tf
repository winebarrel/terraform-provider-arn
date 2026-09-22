# arn:aws:config:ap-northeast-1:111111111111:aggregation-authorization/aggregator-account/aggregator-region
output "config_aggregation_authorization" {
  value = provider::arn::config_aggregation_authorization("aggregator-account", "aggregator-region")
}
