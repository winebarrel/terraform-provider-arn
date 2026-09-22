# arn:aws:athena:ap-northeast-1:111111111111:workgroup/work-group-name/session/session-id
output "athena_session" {
  value = provider::arn::athena_session("work-group-name", "session-id")
}
