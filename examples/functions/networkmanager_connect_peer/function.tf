# arn:aws:networkmanager::111111111111:connect-peer/resource-id
output "networkmanager_connect_peer" {
  value = provider::arn::networkmanager_connect_peer("resource-id")
}
