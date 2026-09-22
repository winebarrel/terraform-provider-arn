# arn:aws:iam::111111111111:user/aws-user-name
output "codestar_user" {
  value = provider::arn::codestar_user("aws-user-name")
}
