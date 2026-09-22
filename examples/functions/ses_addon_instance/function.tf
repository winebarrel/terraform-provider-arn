# arn:aws:ses:ap-northeast-1:111111111111:addon-instance/addon-instance-id
output "ses_addon_instance" {
  value = provider::arn::ses_addon_instance("addon-instance-id")
}
