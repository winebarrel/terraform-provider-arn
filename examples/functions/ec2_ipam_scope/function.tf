# arn:aws:ec2::111111111111:ipam-scope/ipam-scope-id
output "ec2_ipam_scope" {
  value = provider::arn::ec2_ipam_scope("ipam-scope-id")
}
