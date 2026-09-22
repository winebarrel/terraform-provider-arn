# arn:aws:chime:ap-northeast-1:111111111111:app-instance/app-instance-id
output "chime_app_instance" {
  value = provider::arn::chime_app_instance("app-instance-id")
}
