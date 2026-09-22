# arn:aws:ssm-contacts:ap-northeast-1:111111111111:contactchannel/contact-alias/contact-channel-id
output "ssm_contacts_contactchannel" {
  value = provider::arn::ssm_contacts_contactchannel("contact-alias", "contact-channel-id")
}
