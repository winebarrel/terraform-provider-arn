# arn:aws:partnercentral:ap-northeast-1::catalog/catalog/benefit/identifier
output "partnercentral_benefit" {
  value = provider::arn::partnercentral_benefit("catalog", "identifier")
}
