# arn:aws:groundstation:ap-northeast-1:111111111111:contact/contact-id
output "groundstation_contact" {
  value = provider::arn::groundstation_contact("contact-id")
}
