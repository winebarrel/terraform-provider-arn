# arn:aws:partnercentral:ap-northeast-1:111111111111:catalog/catalog/benefit-allocation/identifier
output "partnercentral_benefit_allocation" {
  value = provider::arn::partnercentral_benefit_allocation("catalog", "identifier")
}
