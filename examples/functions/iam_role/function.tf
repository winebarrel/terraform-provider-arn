# arn:aws:iam::111111111111:role/my-role
output "role" {
  value = provider::arn::iam_role("my-role")
}

# arn:aws:iam::222222222222:role/my-role
output "role_in_prod" {
  value = provider::arn::iam_role("my-role", { account = "prod" })
}
