# arn:aws:glue:ap-northeast-1:111111111111:session/session-id
output "glue_session" {
  value = provider::arn::glue_session("session-id")
}
