# arn:aws:ec2:ap-northeast-1:111111111111:vpc-peering-connection/vpc-peering-connection-id
output "ec2_vpc_peering_connection" {
  value = provider::arn::ec2_vpc_peering_connection("vpc-peering-connection-id")
}
