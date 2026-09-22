# arn:aws:partnercentral:ap-northeast-1:111111111111:catalog/catalog/connection-preferences
output "partnercentral_connection_preferences" {
  value = provider::arn::partnercentral_connection_preferences("catalog")
}
