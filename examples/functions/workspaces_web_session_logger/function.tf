# arn:aws:workspaces-web:ap-northeast-1:111111111111:sessionLogger/session-logger-id
output "workspaces_web_session_logger" {
  value = provider::arn::workspaces_web_session_logger("session-logger-id")
}
