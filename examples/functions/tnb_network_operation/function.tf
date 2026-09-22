# arn:aws:tnb:ap-northeast-1:111111111111:network-operation/network-operation-id
output "tnb_network_operation" {
  value = provider::arn::tnb_network_operation("network-operation-id")
}
