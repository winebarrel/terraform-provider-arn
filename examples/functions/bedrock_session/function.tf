# arn:aws:bedrock:ap-northeast-1:111111111111:session/session-id
output "bedrock_session" {
  value = provider::arn::bedrock_session("session-id")
}
