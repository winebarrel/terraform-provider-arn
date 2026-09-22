# arn:aws:ec2:ap-northeast-1:111111111111:ipv4pool-ec2/ipv4-pool-ec2-id
output "ec2_ipv4pool_ec2" {
  value = provider::arn::ec2_ipv4pool_ec2("ipv4-pool-ec2-id")
}
