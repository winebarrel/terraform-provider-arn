# arn:aws:ec2:ap-northeast-1:111111111111:transit-gateway-metering-policy/transit-gateway-metering-policy-id
output "ec2_transit_gateway_metering_policy" {
  value = provider::arn::ec2_transit_gateway_metering_policy("transit-gateway-metering-policy-id")
}
