# arn:aws:ec2:ap-northeast-1:111111111111:vpc-endpoint/vpc-endpoint-id
output "vpce_vpc_endpoint" {
  value = provider::arn::vpce_vpc_endpoint("vpc-endpoint-id")
}
