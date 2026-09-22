# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/campaigns/campaign-id
output "mobiletargeting_campaign" {
  value = provider::arn::mobiletargeting_campaign("app-id", "campaign-id")
}
