# arn:aws:imagebuilder:ap-northeast-1:111111111111:lifecycle-policy/lifecycle-policy-name
output "imagebuilder_lifecycle_policy" {
  value = provider::arn::imagebuilder_lifecycle_policy("lifecycle-policy-name")
}
