# arn:aws:quicksight:ap-northeast-1:111111111111:dashboard/dashboard-id/snapshot-job/resource-id
output "quicksight_dashboard_snapshot_job" {
  value = provider::arn::quicksight_dashboard_snapshot_job("dashboard-id", "resource-id")
}
