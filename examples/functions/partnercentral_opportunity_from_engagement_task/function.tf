# arn:aws:partnercentral:ap-northeast-1::catalog/catalog/opportunity-from-engagement-task/task-id
output "partnercentral_opportunity_from_engagement_task" {
  value = provider::arn::partnercentral_opportunity_from_engagement_task("catalog", "task-id")
}
