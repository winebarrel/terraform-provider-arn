# arn:aws:partnercentral:ap-northeast-1::catalog/catalog/engagement-by-accepting-invitation-task/task-id
output "partnercentral_engagement_by_accepting_invitation_task" {
  value = provider::arn::partnercentral_engagement_by_accepting_invitation_task("catalog", "task-id")
}
