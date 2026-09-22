# arn:aws:config:ap-northeast-1:111111111111:config-aggregator/aggregator-id
output "config_configuration_aggregator" {
  value = provider::arn::config_configuration_aggregator("aggregator-id")
}
