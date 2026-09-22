# arn:aws:snow-device-management:ap-northeast-1:111111111111:managed-device/resource-id
output "snow_device_management_managed_device" {
  value = provider::arn::snow_device_management_managed_device("resource-id")
}
