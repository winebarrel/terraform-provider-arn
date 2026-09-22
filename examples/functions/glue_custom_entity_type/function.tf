# arn:aws:glue:ap-northeast-1:111111111111:customEntityType/custom-entity-type-id
output "glue_custom_entity_type" {
  value = provider::arn::glue_custom_entity_type("custom-entity-type-id")
}
