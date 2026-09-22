# arn:aws:config:ap-northeast-1:111111111111:organization-config-rule/organization-config-rule-id
output "config_organization_config_rule" {
  value = provider::arn::config_organization_config_rule("organization-config-rule-id")
}
