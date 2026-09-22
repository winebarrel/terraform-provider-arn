# arn:aws:ssm:ap-northeast-1:111111111111:maintenancewindow/resource-id
output "ssm_maintenancewindow" {
  value = provider::arn::ssm_maintenancewindow("resource-id")
}
