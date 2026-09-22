# arn:aws:partnercentral:ap-northeast-1:111111111111:catalog/catalog/revenue-attribution/revenue-attribution-id
output "partnercentral_revenue_attribution" {
  value = provider::arn::partnercentral_revenue_attribution("catalog", "revenue-attribution-id")
}
