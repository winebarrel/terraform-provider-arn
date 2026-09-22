# arn:aws:securityhub:ap-northeast-1:111111111111:configuration-policy/configuration-policy-id
output "securityhub_configuration_policy" {
  value = provider::arn::securityhub_configuration_policy("configuration-policy-id")
}
