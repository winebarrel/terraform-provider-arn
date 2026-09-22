# arn:aws:iam::111111111111:role/role-name-with-path
output "ec2_role" {
  value = provider::arn::ec2_role("role-name-with-path")
}
