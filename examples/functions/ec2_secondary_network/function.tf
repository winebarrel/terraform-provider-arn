# arn:aws:ec2:ap-northeast-1:111111111111:secondary-network/secondary-network-id
output "ec2_secondary_network" {
  value = provider::arn::ec2_secondary_network("secondary-network-id")
}
