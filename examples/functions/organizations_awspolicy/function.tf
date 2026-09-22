# arn:aws:organizations::aws:policy/policy-type/p-policy-id
output "organizations_awspolicy" {
  value = provider::arn::organizations_awspolicy("policy-type", "policy-id")
}
