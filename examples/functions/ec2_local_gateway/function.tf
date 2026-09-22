# arn:aws:ec2:ap-northeast-1:111111111111:local-gateway/local-gateway-id
output "ec2_local_gateway" {
  value = provider::arn::ec2_local_gateway("local-gateway-id")
}
