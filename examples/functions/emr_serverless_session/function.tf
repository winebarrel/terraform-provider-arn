# arn:aws:emr-serverless:ap-northeast-1:111111111111:/applications/application-id/sessions/session-id
output "emr_serverless_session" {
  value = provider::arn::emr_serverless_session("application-id", "session-id")
}
