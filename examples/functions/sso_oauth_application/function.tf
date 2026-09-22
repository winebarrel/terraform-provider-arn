# arn:aws:sso::111111111111:application/instance-id/application-id
output "sso_oauth_application" {
  value = provider::arn::sso_oauth_application("instance-id", "application-id")
}
