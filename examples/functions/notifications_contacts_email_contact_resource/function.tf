# arn:aws:notifications-contacts::111111111111:emailcontact/email-contact-id
output "notifications_contacts_email_contact_resource" {
  value = provider::arn::notifications_contacts_email_contact_resource("email-contact-id")
}
