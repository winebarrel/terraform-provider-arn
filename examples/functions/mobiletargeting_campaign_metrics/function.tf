# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/campaigns/campaign-id/kpis/daterange/kpi-name
output "mobiletargeting_campaign_metrics" {
  value = provider::arn::mobiletargeting_campaign_metrics("app-id", "campaign-id", "kpi-name")
}
