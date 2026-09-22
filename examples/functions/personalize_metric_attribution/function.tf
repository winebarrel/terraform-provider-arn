# arn:aws:personalize:ap-northeast-1:111111111111:metric-attribution/resource-id
output "personalize_metric_attribution" {
  value = provider::arn::personalize_metric_attribution("resource-id")
}
