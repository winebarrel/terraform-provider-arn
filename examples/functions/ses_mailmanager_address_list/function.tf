# arn:aws:ses:ap-northeast-1:111111111111:mailmanager-address-list/address-list-id
output "ses_mailmanager_address_list" {
  value = provider::arn::ses_mailmanager_address_list("address-list-id")
}
