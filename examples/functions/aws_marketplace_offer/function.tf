# arn:aws:aws-marketplace:::catalog/catalog-name/offer/offer-id
output "aws_marketplace_offer" {
  value = provider::arn::aws_marketplace_offer("catalog-name", "offer-id")
}
