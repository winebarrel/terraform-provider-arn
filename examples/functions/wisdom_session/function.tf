# arn:aws:wisdom:ap-northeast-1:111111111111:session/assistant-id/session-id
output "wisdom_session" {
  value = provider::arn::wisdom_session("assistant-id", "session-id")
}
