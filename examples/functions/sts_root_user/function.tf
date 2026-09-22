# arn:aws:iam::111111111111:root
output "sts_root_user" {
  value = provider::arn::sts_root_user()
}
