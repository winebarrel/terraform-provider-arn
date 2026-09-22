# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/journeys
output "mobiletargeting_journeys" {
  value = provider::arn::mobiletargeting_journeys("app-id")
}
