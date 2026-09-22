# arn:aws:aws-marketplace:ap-northeast-1::catalog/Assessment/resource-id
output "aws_marketplace_assessment" {
  value = provider::arn::aws_marketplace_assessment("catalog", "resource-id")
}
