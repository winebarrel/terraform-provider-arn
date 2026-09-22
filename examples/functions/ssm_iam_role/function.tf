# arn:aws:iam::111111111111:role/role-name
output "ssm_iam_role" {
  value = provider::arn::ssm_iam_role("role-name")
}
