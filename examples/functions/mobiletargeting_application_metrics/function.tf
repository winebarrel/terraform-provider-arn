# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/kpis/daterange/kpi-name
output "mobiletargeting_application_metrics" {
  value = provider::arn::mobiletargeting_application_metrics("app-id", "kpi-name")
}
