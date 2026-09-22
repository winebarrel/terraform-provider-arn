# arn:aws:sagemaker:ap-northeast-1:111111111111:device-fleet/device-fleet-name/device/device-name
output "sagemaker_device" {
  value = provider::arn::sagemaker_device("device-fleet-name", "device-name")
}
