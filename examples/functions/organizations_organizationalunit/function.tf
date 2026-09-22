# arn:aws:organizations::111111111111:ou/o-organization-id/ou-organizational-unit-id
output "organizations_organizationalunit" {
  value = provider::arn::organizations_organizationalunit("organization-id", "organizational-unit-id")
}
