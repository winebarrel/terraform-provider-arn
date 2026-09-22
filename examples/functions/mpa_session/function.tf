# arn:aws:mpa:ap-northeast-1:111111111111:session/session-id
output "mpa_session" {
  value = provider::arn::mpa_session("session-id")
}
