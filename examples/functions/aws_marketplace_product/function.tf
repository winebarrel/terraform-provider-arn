# arn:aws:aws-marketplace:::catalog/catalog-name/product/product-id
output "aws_marketplace_product" {
  value = provider::arn::aws_marketplace_product("catalog-name", "product-id")
}
