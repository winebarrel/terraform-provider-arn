# arn:aws:ec2:ap-northeast-1:111111111111:transit-gateway-connect-peer/transit-gateway-connect-peer-id
output "ec2_transit_gateway_connect_peer" {
  value = provider::arn::ec2_transit_gateway_connect_peer("transit-gateway-connect-peer-id")
}
