# arn:aws:panorama:ap-northeast-1:111111111111:device/device-id
output "panorama_device" {
  value = provider::arn::panorama_device("device-id")
}
