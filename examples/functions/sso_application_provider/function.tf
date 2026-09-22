# arn:aws:sso::aws:applicationProvider/application-provider-id
output "sso_application_provider" {
  value = provider::arn::sso_application_provider("application-provider-id")
}
