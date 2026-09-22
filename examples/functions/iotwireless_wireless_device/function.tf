# arn:aws:iotwireless:ap-northeast-1:111111111111:WirelessDevice/wireless-device-id
output "iotwireless_wireless_device" {
  value = provider::arn::iotwireless_wireless_device("wireless-device-id")
}
