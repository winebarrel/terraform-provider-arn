# arn:aws:transform-custom:ap-northeast-1:111111111111:finding/finding-id
output "transform_custom_finding" {
  value = provider::arn::transform_custom_finding("finding-id")
}
