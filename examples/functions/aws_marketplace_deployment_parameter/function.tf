# arn:aws:aws-marketplace:ap-northeast-1:111111111111:DeploymentParameter:catalogs/catalog-name/products/product-id/resource-id
output "aws_marketplace_deployment_parameter" {
  value = provider::arn::aws_marketplace_deployment_parameter("catalog-name", "product-id", "resource-id")
}
