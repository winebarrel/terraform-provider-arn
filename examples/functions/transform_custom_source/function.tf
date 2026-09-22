# arn:aws:transform-custom:ap-northeast-1:111111111111:source/name
output "transform_custom_source" {
  value = provider::arn::transform_custom_source("name")
}
