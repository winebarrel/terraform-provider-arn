# arn:aws:personalize:ap-northeast-1:111111111111:data-insights-job/resource-id
output "personalize_data_insights_job" {
  value = provider::arn::personalize_data_insights_job("resource-id")
}
