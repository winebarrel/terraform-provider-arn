# arn:aws:ec2:ap-northeast-1:111111111111:transit-gateway-multicast-domain/transit-gateway-multicast-domain-id
output "ec2_transit_gateway_multicast_domain" {
  value = provider::arn::ec2_transit_gateway_multicast_domain("transit-gateway-multicast-domain-id")
}
