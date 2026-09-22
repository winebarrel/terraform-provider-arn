# arn:aws:iam::111111111111:role/role-name-with-path
output "sts_role" {
  value = provider::arn::sts_role("role-name-with-path")
}
