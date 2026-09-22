# arn:aws:organizations::111111111111:policy/o-organization-id/policy-type/p-policy-id
output "organizations_policy" {
  value = provider::arn::organizations_policy("organization-id", "policy-type", "policy-id")
}
