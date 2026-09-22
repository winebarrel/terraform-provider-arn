# arn:aws:networkmanager::111111111111:site/global-network-id/resource-id
output "networkmanager_site" {
  value = provider::arn::networkmanager_site("global-network-id", "resource-id")
}
