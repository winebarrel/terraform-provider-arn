# arn:aws:ssm-contacts:ap-northeast-1:111111111111:page/contact-alias/page-id
output "ssm_contacts_page" {
  value = provider::arn::ssm_contacts_page("contact-alias", "page-id")
}
