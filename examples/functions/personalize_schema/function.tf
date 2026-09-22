# arn:aws:personalize:ap-northeast-1:111111111111:schema/resource-id
output "personalize_schema" {
  value = provider::arn::personalize_schema("resource-id")
}
