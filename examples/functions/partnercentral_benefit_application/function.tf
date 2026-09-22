# arn:aws:partnercentral:ap-northeast-1:111111111111:catalog/catalog/benefit-application/identifier
output "partnercentral_benefit_application" {
  value = provider::arn::partnercentral_benefit_application("catalog", "identifier")
}
