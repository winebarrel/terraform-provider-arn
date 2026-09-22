# arn:aws:partnercentral:ap-northeast-1::catalog/catalog/connection-invitation/identifier
output "partnercentral_connection_invitation" {
  value = provider::arn::partnercentral_connection_invitation("catalog", "identifier")
}
