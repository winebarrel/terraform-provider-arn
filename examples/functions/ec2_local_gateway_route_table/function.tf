# arn:aws:ec2:ap-northeast-1:111111111111:local-gateway-route-table/local-gateway-routetable-id
output "ec2_local_gateway_route_table" {
  value = provider::arn::ec2_local_gateway_route_table("local-gateway-routetable-id")
}
