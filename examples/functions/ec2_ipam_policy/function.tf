# arn:aws:ec2::111111111111:ipam-policy/ipam-policy-id
output "ec2_ipam_policy" {
  value = provider::arn::ec2_ipam_policy("ipam-policy-id")
}
