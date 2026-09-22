# arn:aws:networkmanager::111111111111:global-network/resource-id
output "networkmanager_global_network" {
  value = provider::arn::networkmanager_global_network("resource-id")
}
