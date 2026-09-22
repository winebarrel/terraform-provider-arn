# arn:aws:connect-campaigns:ap-northeast-1:111111111111:campaign/campaign-id
output "connect_campaigns_campaign" {
  value = provider::arn::connect_campaigns_campaign("campaign-id")
}
