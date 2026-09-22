# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/journeys/journey-id
output "mobiletargeting_journey" {
  value = provider::arn::mobiletargeting_journey("app-id", "journey-id")
}
