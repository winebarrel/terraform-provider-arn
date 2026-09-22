# arn:aws:snow-device-management:ap-northeast-1:111111111111:task/resource-id
output "snow_device_management_task" {
  value = provider::arn::snow_device_management_task("resource-id")
}
