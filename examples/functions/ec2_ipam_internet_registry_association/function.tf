# arn:aws:ec2::111111111111:ipam-internet-registry-association/ipam-internet-registry-association-id
output "ec2_ipam_internet_registry_association" {
  value = provider::arn::ec2_ipam_internet_registry_association("ipam-internet-registry-association-id")
}
