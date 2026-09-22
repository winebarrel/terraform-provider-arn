# arn:aws:appstudio:ap-northeast-1:111111111111:instance/instance-id
output "appstudio_instance" {
  value = provider::arn::appstudio_instance("instance-id")
}
