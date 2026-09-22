# arn:aws:ec2:ap-northeast-1:111111111111:transit-gateway-attachment/transit-gateway-attachment-id
output "ec2_transit_gateway_attachment" {
  value = provider::arn::ec2_transit_gateway_attachment("transit-gateway-attachment-id")
}
