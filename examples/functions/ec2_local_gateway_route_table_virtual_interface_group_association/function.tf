# arn:aws:ec2:ap-northeast-1:111111111111:local-gateway-route-table-virtual-interface-group-association/local-gateway-route-table-virtual-interface-group-association-id
output "ec2_local_gateway_route_table_virtual_interface_group_association" {
  value = provider::arn::ec2_local_gateway_route_table_virtual_interface_group_association("local-gateway-route-table-virtual-interface-group-association-id")
}
