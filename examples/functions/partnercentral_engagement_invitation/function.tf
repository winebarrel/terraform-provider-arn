# arn:aws:partnercentral:ap-northeast-1::catalog/catalog/engagement-invitation/identifier
output "partnercentral_engagement_invitation" {
  value = provider::arn::partnercentral_engagement_invitation("catalog", "identifier")
}
