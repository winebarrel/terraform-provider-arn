# arn:aws:workspaces-web:ap-northeast-1:111111111111:dataProtectionSettings/data-protection-settings-id
output "workspaces_web_data_protection_settings" {
  value = provider::arn::workspaces_web_data_protection_settings("data-protection-settings-id")
}
