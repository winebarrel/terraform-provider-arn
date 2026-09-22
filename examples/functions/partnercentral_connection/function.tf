# arn:aws:partnercentral:ap-northeast-1::catalog/catalog/connection/identifier
output "partnercentral_connection" {
  value = provider::arn::partnercentral_connection("catalog", "identifier")
}
