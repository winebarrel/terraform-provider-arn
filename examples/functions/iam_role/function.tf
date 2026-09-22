# arn:aws:iam::111111111111:role/role-name-with-path
output "iam_role" {
  value = provider::arn::iam_role("role-name-with-path")
}
