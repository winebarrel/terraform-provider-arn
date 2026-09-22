# arn:aws:aws-marketplace:::catalog/catalog-name/purchaseOption/purchase-option-id
output "aws_marketplace_purchase_option" {
  value = provider::arn::aws_marketplace_purchase_option("catalog-name", "purchase-option-id")
}
