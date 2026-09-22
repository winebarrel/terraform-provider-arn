# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/journeys/journey-id/kpis/daterange/kpi-name
output "mobiletargeting_journey_metrics" {
  value = provider::arn::mobiletargeting_journey_metrics("app-id", "journey-id", "kpi-name")
}
