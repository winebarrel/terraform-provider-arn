# arn:aws:qbusiness:ap-northeast-1:111111111111:application/application-id
output "qapps_application" {
  value = provider::arn::qapps_application("application-id")
}
