# arn:aws:partnercentral:ap-northeast-1:111111111111:catalog/catalog/partner/identifier
output "partnercentral_partner" {
  value = provider::arn::partnercentral_partner("catalog", "identifier")
}
