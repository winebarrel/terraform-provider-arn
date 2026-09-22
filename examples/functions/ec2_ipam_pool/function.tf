# arn:aws:ec2::111111111111:ipam-pool/ipam-pool-id
output "ec2_ipam_pool" {
  value = provider::arn::ec2_ipam_pool("ipam-pool-id")
}
