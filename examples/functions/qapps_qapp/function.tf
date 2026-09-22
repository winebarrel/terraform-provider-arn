# arn:aws:qapps:ap-northeast-1:111111111111:application/application-id/qapp/app-id
output "qapps_qapp" {
  value = provider::arn::qapps_qapp("application-id", "app-id")
}
