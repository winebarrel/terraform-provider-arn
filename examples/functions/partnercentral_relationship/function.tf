# arn:aws:partnercentral:ap-northeast-1:111111111111:catalog/catalog/program-management-account/program-management-account-id/relationship/relationship-id
output "partnercentral_relationship" {
  value = provider::arn::partnercentral_relationship("catalog", "program-management-account-id", "relationship-id")
}
