# arn:aws:organizations::111111111111:resourcepolicy/o-organization-id/rp-resource-policy-id
output "organizations_resourcepolicy" {
  value = provider::arn::organizations_resourcepolicy("organization-id", "resource-policy-id")
}
