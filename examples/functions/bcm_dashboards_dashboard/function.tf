# arn:aws:bcm-dashboards::111111111111:dashboard/dashboard-name
output "bcm_dashboards_dashboard" {
  value = provider::arn::bcm_dashboards_dashboard("dashboard-name")
}
