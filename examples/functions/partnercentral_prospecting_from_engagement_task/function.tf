# arn:aws:partnercentral:ap-northeast-1::catalog/catalog/prospecting-from-engagement-task/task-identifier
output "partnercentral_prospecting_from_engagement_task" {
  value = provider::arn::partnercentral_prospecting_from_engagement_task("catalog", "task-identifier")
}
