# arn:aws:ec2::111111111111:ipam/ipam-id
output "ec2_ipam" {
  value = provider::arn::ec2_ipam("ipam-id")
}
