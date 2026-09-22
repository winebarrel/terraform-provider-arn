# arn:aws:workspaces-web:ap-northeast-1:111111111111:userAccessLoggingSettings/user-access-logging-settings-id
output "workspaces_web_user_access_logging_settings" {
  value = provider::arn::workspaces_web_user_access_logging_settings("user-access-logging-settings-id")
}
