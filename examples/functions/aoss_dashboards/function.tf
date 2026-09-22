# arn:aws:aoss:ap-northeast-1:111111111111:dashboards/default
output "aoss_dashboards" {
  value = provider::arn::aoss_dashboards()
}
