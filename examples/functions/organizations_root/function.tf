# arn:aws:organizations::111111111111:root/o-organization-id/r-root-id
output "organizations_root" {
  value = provider::arn::organizations_root("organization-id", "root-id")
}
