# arn:aws:aws-marketplace:::catalog/catalog-name/offerSet/offer-set-id
output "aws_marketplace_offer_set" {
  value = provider::arn::aws_marketplace_offer_set("catalog-name", "offer-set-id")
}
