# arn:aws:macie2:ap-northeast-1:111111111111:custom-data-identifier/resource-id
output "macie2_custom_data_identifier" {
  value = provider::arn::macie2_custom_data_identifier("resource-id")
}
