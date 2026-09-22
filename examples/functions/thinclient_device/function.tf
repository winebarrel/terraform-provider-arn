# arn:aws:thinclient:ap-northeast-1:111111111111:device/device-id
output "thinclient_device" {
  value = provider::arn::thinclient_device("device-id")
}
