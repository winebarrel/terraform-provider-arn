# arn:aws:partnercentral:ap-northeast-1::catalog/catalog/engagement-from-opportunity-task/task-id
output "partnercentral_engagement_from_opportunity_task" {
  value = provider::arn::partnercentral_engagement_from_opportunity_task("catalog", "task-id")
}
