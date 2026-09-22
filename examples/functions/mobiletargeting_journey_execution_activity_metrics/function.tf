# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/journeys/journey-id/activities/journey-activity-id/execution-metrics
output "mobiletargeting_journey_execution_activity_metrics" {
  value = provider::arn::mobiletargeting_journey_execution_activity_metrics("app-id", "journey-id", "journey-activity-id")
}
