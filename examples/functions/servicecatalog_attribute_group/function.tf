# arn:aws:servicecatalog:ap-northeast-1:111111111111:/attribute-groups/attribute-group-id
output "servicecatalog_attribute_group" {
  value = provider::arn::servicecatalog_attribute_group("attribute-group-id")
}
