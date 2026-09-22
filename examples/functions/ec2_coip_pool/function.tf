# arn:aws:ec2:ap-northeast-1:111111111111:coip-pool/ipv4-pool-coip-id
output "ec2_coip_pool" {
  value = provider::arn::ec2_coip_pool("ipv4-pool-coip-id")
}
