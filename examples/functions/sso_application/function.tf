# arn:aws:sso::111111111111:application/instance-id/application-id
output "sso_application" {
  value = provider::arn::sso_application("instance-id", "application-id")
}
