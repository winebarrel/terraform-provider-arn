# arn:aws:ec2:ap-northeast-1:111111111111:ipv6pool-ec2/ipv6-pool-ec2-id
output "ec2_ipv6pool_ec2" {
  value = provider::arn::ec2_ipv6pool_ec2("ipv6-pool-ec2-id")
}
