# arn:aws:ec2:ap-northeast-1:111111111111:local-gateway-route-table-vpc-association/local-gateway-route-table-vpc-association-id
output "ec2_local_gateway_route_table_vpc_association" {
  value = provider::arn::ec2_local_gateway_route_table_vpc_association("local-gateway-route-table-vpc-association-id")
}
