# arn:aws:networkmanager::111111111111:device/global-network-id/resource-id
output "networkmanager_device" {
  value = provider::arn::networkmanager_device("global-network-id", "resource-id")
}
