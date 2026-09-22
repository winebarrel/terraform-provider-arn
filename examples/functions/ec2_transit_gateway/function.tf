# arn:aws:ec2:ap-northeast-1:111111111111:transit-gateway/transit-gateway-id
output "ec2_transit_gateway" {
  value = provider::arn::ec2_transit_gateway("transit-gateway-id")
}
