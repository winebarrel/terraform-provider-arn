# arn:aws:ec2::111111111111:ipam-prefix-list-resolver-target/ipam-prefix-list-resolver-target-id
output "ec2_ipam_prefix_list_resolver_target" {
  value = provider::arn::ec2_ipam_prefix_list_resolver_target("ipam-prefix-list-resolver-target-id")
}
