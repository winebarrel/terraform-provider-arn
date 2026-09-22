# arn:aws:ec2::111111111111:ipam-resource-discovery-association/ipam-resource-discovery-association-id
output "ec2_ipam_resource_discovery_association" {
  value = provider::arn::ec2_ipam_resource_discovery_association("ipam-resource-discovery-association-id")
}
