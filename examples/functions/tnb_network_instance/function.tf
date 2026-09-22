# arn:aws:tnb:ap-northeast-1:111111111111:network-instance/network-instance-id
output "tnb_network_instance" {
  value = provider::arn::tnb_network_instance("network-instance-id")
}
