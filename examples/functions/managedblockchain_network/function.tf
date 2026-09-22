# arn:aws:managedblockchain:ap-northeast-1::networks/network-id
output "managedblockchain_network" {
  value = provider::arn::managedblockchain_network("network-id")
}
