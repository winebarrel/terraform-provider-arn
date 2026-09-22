# arn:aws:aws-marketplace:ap-northeast-1:111111111111:catalog/entity-type/resource-id
output "aws_marketplace_entity" {
  value = provider::arn::aws_marketplace_entity("catalog", "entity-type", "resource-id")
}
