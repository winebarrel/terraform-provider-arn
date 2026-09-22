# arn:aws:partnercentral:ap-northeast-1:111111111111:catalog/catalog/marketplace-revenue-share/marketplace-product-id
output "partnercentral_marketplace_revenue_share" {
  value = provider::arn::partnercentral_marketplace_revenue_share("catalog", "marketplace-product-id")
}
