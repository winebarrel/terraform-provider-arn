# arn:aws:transform:ap-northeast-1:111111111111:profile/identifier
output "transform_profile" {
  value = provider::arn::transform_profile("identifier")
}
