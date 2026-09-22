# arn:aws:cleanrooms:ap-northeast-1:111111111111:membership/membership-id/privacybudgettemplate/privacy-budget-template-id
output "cleanrooms_privacybudgettemplate" {
  value = provider::arn::cleanrooms_privacybudgettemplate("membership-id", "privacy-budget-template-id")
}
