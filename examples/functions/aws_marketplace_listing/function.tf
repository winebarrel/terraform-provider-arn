# arn:aws:aws-marketplace:::catalog/catalog-name/listing/listing-id
output "aws_marketplace_listing" {
  value = provider::arn::aws_marketplace_listing("catalog-name", "listing-id")
}
