# arn:aws:partnercentral:ap-northeast-1:111111111111:catalog/catalog/solution/identifier
output "partnercentral_solution" {
  value = provider::arn::partnercentral_solution("catalog", "identifier")
}
