# arn:aws:comprehend:ap-northeast-1:111111111111:entity-recognizer-endpoint/entity-recognizer-endpoint-name
output "comprehend_entity_recognizer_endpoint" {
  value = provider::arn::comprehend_entity_recognizer_endpoint("entity-recognizer-endpoint-name")
}
