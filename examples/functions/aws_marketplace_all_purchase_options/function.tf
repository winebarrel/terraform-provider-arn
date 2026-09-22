# arn:aws:aws-marketplace:::catalog/catalog-name/purchaseOption/*
output "aws_marketplace_all_purchase_options" {
  value = provider::arn::aws_marketplace_all_purchase_options("catalog-name")
}
