# arn:aws:appstudio:ap-northeast-1:111111111111:application/application-id
output "appstudio_application" {
  value = provider::arn::appstudio_application("application-id")
}
