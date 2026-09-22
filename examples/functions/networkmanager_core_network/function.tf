# arn:aws:networkmanager::111111111111:core-network/resource-id
output "networkmanager_core_network" {
  value = provider::arn::networkmanager_core_network("resource-id")
}
