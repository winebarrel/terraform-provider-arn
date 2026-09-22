# arn:aws:partnercentral:ap-northeast-1:111111111111:catalog/catalog/opportunity/identifier
output "partnercentral_opportunity" {
  value = provider::arn::partnercentral_opportunity("catalog", "identifier")
}
