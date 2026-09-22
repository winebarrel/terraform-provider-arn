# arn:aws:organizations::111111111111:organization/o-organization-id
output "organizations_organization" {
  value = provider::arn::organizations_organization("organization-id")
}
