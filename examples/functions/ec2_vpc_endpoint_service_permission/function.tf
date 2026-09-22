# arn:aws:ec2:ap-northeast-1:111111111111:vpc-endpoint-service-permission/vpc-endpoint-service-permission-id
output "ec2_vpc_endpoint_service_permission" {
  value = provider::arn::ec2_vpc_endpoint_service_permission("vpc-endpoint-service-permission-id")
}
