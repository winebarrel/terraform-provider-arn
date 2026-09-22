# arn:aws:ec2:ap-northeast-1:111111111111:transit-gateway-route-table/transit-gateway-route-table-id
output "ec2_transit_gateway_route_table" {
  value = provider::arn::ec2_transit_gateway_route_table("transit-gateway-route-table-id")
}
