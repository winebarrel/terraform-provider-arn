# arn:aws:ssm-contacts:ap-northeast-1:111111111111:engagement/contact-alias/engagement-id
output "ssm_contacts_engagement" {
  value = provider::arn::ssm_contacts_engagement("contact-alias", "engagement-id")
}
