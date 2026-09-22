# arn:aws:appstream:ap-northeast-1:111111111111:application/application-name
output "appstream_application" {
  value = provider::arn::appstream_application("application-name")
}
