# arn:aws:wickr:ap-northeast-1:111111111111:network/network-id
output "wickr_network" {
  value = provider::arn::wickr_network("network-id")
}
