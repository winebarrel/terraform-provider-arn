# arn:aws:aws-marketplace:ap-northeast-1:111111111111:catalog/ChangeSet/resource-id
output "aws_marketplace_change_set" {
  value = provider::arn::aws_marketplace_change_set("catalog", "resource-id")
}
