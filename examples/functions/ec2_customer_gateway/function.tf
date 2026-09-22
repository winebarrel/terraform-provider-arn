# arn:aws:ec2:ap-northeast-1:111111111111:customer-gateway/customer-gateway-id
output "ec2_customer_gateway" {
  value = provider::arn::ec2_customer_gateway("customer-gateway-id")
}
