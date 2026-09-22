# arn:aws:networkmanager::111111111111:connection/global-network-id/resource-id
output "networkmanager_connection" {
  value = provider::arn::networkmanager_connection("global-network-id", "resource-id")
}
