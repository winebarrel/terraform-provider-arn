# arn:aws:config:ap-northeast-1:111111111111:config-rule/config-rule-id
output "config_config_rule" {
  value = provider::arn::config_config_rule("config-rule-id")
}
