# arn:aws:frauddetector:ap-northeast-1:111111111111:entity-type/resource-path
output "frauddetector_entity_type" {
  value = provider::arn::frauddetector_entity_type("resource-path")
}
