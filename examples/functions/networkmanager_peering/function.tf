# arn:aws:networkmanager::111111111111:peering/resource-id
output "networkmanager_peering" {
  value = provider::arn::networkmanager_peering("resource-id")
}
