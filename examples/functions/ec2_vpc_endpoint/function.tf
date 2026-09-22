# arn:aws:ec2:ap-northeast-1:111111111111:vpc-endpoint/vpc-endpoint-id
output "ec2_vpc_endpoint" {
  value = provider::arn::ec2_vpc_endpoint("vpc-endpoint-id")
}
