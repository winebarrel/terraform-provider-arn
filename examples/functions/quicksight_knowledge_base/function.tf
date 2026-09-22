# arn:aws:quicksight:ap-northeast-1:111111111111:knowledge-base/resource-id
output "quicksight_knowledge_base" {
  value = provider::arn::quicksight_knowledge_base("resource-id")
}
