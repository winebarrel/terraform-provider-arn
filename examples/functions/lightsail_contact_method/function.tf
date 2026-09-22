# arn:aws:lightsail:ap-northeast-1:111111111111:ContactMethod/id
output "lightsail_contact_method" {
  value = provider::arn::lightsail_contact_method("id")
}
