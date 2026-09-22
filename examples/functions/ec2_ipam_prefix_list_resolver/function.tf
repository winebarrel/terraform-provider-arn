# arn:aws:ec2::111111111111:ipam-prefix-list-resolver/ipam-prefix-list-resolver-id
output "ec2_ipam_prefix_list_resolver" {
  value = provider::arn::ec2_ipam_prefix_list_resolver("ipam-prefix-list-resolver-id")
}
