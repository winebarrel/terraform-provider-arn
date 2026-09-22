# arn:aws:bcm-dashboards::111111111111:scheduled-report/scheduled-report-name
output "bcm_dashboards_scheduled_report" {
  value = provider::arn::bcm_dashboards_scheduled_report("scheduled-report-name")
}
