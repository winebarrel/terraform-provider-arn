# arn:aws:iam::111111111111:assumed-role/role-name/role-session-name
output "iam_assumed_role" {
  value = provider::arn::iam_assumed_role("role-name", "role-session-name")
}
