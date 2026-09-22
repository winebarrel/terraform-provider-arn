# arn:aws:ec2::111111111111:ipam-resource-discovery/ipam-resource-discovery-id
output "ec2_ipam_resource_discovery" {
  value = provider::arn::ec2_ipam_resource_discovery("ipam-resource-discovery-id")
}
