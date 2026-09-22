# arn:aws:ec2:ap-northeast-1:111111111111:ipam-pool-allocation/ipam-pool-allocation-id
output "ec2_ipam_pool_allocation" {
  value = provider::arn::ec2_ipam_pool_allocation("ipam-pool-allocation-id")
}
