# arn:aws:transform-custom:ap-northeast-1:111111111111:package/transformation-package-name/knowledge-item/id
output "transform_custom_knowledge_item" {
  value = provider::arn::transform_custom_knowledge_item("transformation-package-name", "id")
}
