# arn:aws:transform-custom:ap-northeast-1:111111111111:package/name
output "transform_custom_package" {
  value = provider::arn::transform_custom_package("name")
}
