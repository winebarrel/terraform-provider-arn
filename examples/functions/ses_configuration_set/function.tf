# arn:aws:ses:ap-northeast-1:111111111111:configuration-set/configuration-set-name
output "ses_configuration_set" {
  value = provider::arn::ses_configuration_set("configuration-set-name")
}
