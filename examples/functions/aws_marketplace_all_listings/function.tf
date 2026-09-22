# arn:aws:aws-marketplace:::catalog/catalog-name/listing/*
output "aws_marketplace_all_listings" {
  value = provider::arn::aws_marketplace_all_listings("catalog-name")
}
