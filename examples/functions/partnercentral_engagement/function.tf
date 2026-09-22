# arn:aws:partnercentral:ap-northeast-1::catalog/catalog/engagement/identifier
output "partnercentral_engagement" {
  value = provider::arn::partnercentral_engagement("catalog", "identifier")
}
