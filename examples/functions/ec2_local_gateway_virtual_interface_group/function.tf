# arn:aws:ec2:ap-northeast-1:111111111111:local-gateway-virtual-interface-group/local-gateway-virtual-interface-group-id
output "ec2_local_gateway_virtual_interface_group" {
  value = provider::arn::ec2_local_gateway_virtual_interface_group("local-gateway-virtual-interface-group-id")
}
