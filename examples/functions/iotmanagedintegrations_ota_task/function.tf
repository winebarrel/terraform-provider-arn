# arn:aws:iotmanagedintegrations:ap-northeast-1:111111111111:ota-task/identifier
output "iotmanagedintegrations_ota_task" {
  value = provider::arn::iotmanagedintegrations_ota_task("identifier")
}
