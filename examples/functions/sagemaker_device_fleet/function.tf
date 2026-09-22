# arn:aws:sagemaker:ap-northeast-1:111111111111:device-fleet/device-fleet-name
output "sagemaker_device_fleet" {
  value = provider::arn::sagemaker_device_fleet("device-fleet-name")
}
