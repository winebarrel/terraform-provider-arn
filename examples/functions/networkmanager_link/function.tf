# arn:aws:networkmanager::111111111111:link/global-network-id/resource-id
output "networkmanager_link" {
  value = provider::arn::networkmanager_link("global-network-id", "resource-id")
}
