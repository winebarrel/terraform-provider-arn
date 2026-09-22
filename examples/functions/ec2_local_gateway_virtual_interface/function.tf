# arn:aws:ec2:ap-northeast-1:111111111111:local-gateway-virtual-interface/local-gateway-virtual-interface-id
output "ec2_local_gateway_virtual_interface" {
  value = provider::arn::ec2_local_gateway_virtual_interface("local-gateway-virtual-interface-id")
}
