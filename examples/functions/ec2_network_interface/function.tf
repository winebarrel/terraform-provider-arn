# arn:aws:ec2:ap-northeast-1:111111111111:network-interface/network-interface-id
output "ec2_network_interface" {
  value = provider::arn::ec2_network_interface("network-interface-id")
}
