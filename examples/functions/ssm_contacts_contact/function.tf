# arn:aws:ssm-contacts:ap-northeast-1:111111111111:contact/contact-alias
output "ssm_contacts_contact" {
  value = provider::arn::ssm_contacts_contact("contact-alias")
}
