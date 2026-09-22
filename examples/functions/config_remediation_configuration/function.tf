# arn:aws:config:ap-northeast-1:111111111111:remediation-configuration/remediation-configuration-id
output "config_remediation_configuration" {
  value = provider::arn::config_remediation_configuration("remediation-configuration-id")
}
