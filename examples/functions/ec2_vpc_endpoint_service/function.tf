# arn:aws:ec2:ap-northeast-1:111111111111:vpc-endpoint-service/vpc-endpoint-service-id
output "ec2_vpc_endpoint_service" {
  value = provider::arn::ec2_vpc_endpoint_service("vpc-endpoint-service-id")
}
