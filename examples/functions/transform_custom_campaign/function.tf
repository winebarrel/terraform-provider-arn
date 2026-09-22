# arn:aws:transform-custom:ap-northeast-1:111111111111:campaign/name
output "transform_custom_campaign" {
  value = provider::arn::transform_custom_campaign("name")
}
