# arn:aws:qapps:ap-northeast-1:111111111111:application/application-id/qapp/app-id/session/session-id
output "qapps_qapp_session" {
  value = provider::arn::qapps_qapp_session("application-id", "app-id", "session-id")
}
