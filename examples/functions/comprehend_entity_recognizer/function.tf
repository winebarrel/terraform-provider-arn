# arn:aws:comprehend:ap-northeast-1:111111111111:entity-recognizer/entity-recognizer-name
output "comprehend_entity_recognizer" {
  value = provider::arn::comprehend_entity_recognizer("entity-recognizer-name")
}
