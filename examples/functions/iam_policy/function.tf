# arn:aws:iam::111111111111:policy/policy-name-with-path
output "iam_policy" {
  value = provider::arn::iam_policy("policy-name-with-path")
}
