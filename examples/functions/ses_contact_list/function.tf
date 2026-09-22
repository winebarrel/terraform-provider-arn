# arn:aws:ses:ap-northeast-1:111111111111:contact-list/contact-list-name
output "ses_contact_list" {
  value = provider::arn::ses_contact_list("contact-list-name")
}
