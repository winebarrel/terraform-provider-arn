# arn:aws:ssm:ap-northeast-1:111111111111:session/session-id
output "ssm_session" {
  value = provider::arn::ssm_session("session-id")
}
