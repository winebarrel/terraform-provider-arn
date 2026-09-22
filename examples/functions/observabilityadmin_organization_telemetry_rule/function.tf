# arn:aws:observabilityadmin:ap-northeast-1:111111111111:organization-telemetry-rule/telemetry-rule-name
output "observabilityadmin_organization_telemetry_rule" {
  value = provider::arn::observabilityadmin_organization_telemetry_rule("telemetry-rule-name")
}
