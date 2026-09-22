# arn:aws:chime:ap-northeast-1:111111111111:app-instance/app-instance-id/user/app-instance-user-id
output "chime_app_instance_user" {
  value = provider::arn::chime_app_instance_user("app-instance-id", "app-instance-user-id")
}
