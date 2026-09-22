# arn:aws:sso:::instance/instance-id
output "sso_instance" {
  value = provider::arn::sso_instance("instance-id")
}
