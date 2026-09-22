# arn:aws:partnercentral:ap-northeast-1:111111111111:catalog/catalog/program-management-account/identifier
output "partnercentral_program_management_account" {
  value = provider::arn::partnercentral_program_management_account("catalog", "identifier")
}
