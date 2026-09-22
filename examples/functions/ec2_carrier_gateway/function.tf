# arn:aws:ec2:ap-northeast-1:111111111111:carrier-gateway/carrier-gateway-id
output "ec2_carrier_gateway" {
  value = provider::arn::ec2_carrier_gateway("carrier-gateway-id")
}
