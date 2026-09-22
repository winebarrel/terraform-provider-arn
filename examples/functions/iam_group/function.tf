# arn:aws:iam::111111111111:group/group-name-with-path
output "iam_group" {
  value = provider::arn::iam_group("group-name-with-path")
}
