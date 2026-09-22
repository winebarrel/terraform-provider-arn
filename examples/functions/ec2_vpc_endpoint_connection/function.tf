# arn:aws:ec2:ap-northeast-1:111111111111:vpc-endpoint-connection/vpc-endpoint-connection-id
output "ec2_vpc_endpoint_connection" {
  value = provider::arn::ec2_vpc_endpoint_connection("vpc-endpoint-connection-id")
}
