# arn:aws:iam::111111111111:user/user-name-with-path
output "iam_user" {
  value = provider::arn::iam_user("user-name-with-path")
}
