# arn:aws:ec2:ap-northeast-1:111111111111:transit-gateway-route-table-announcement/transit-gateway-route-table-announcement-id
output "ec2_transit_gateway_route_table_announcement" {
  value = provider::arn::ec2_transit_gateway_route_table_announcement("transit-gateway-route-table-announcement-id")
}
