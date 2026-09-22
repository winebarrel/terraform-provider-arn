# arn:aws:observabilityadmin:ap-northeast-1:111111111111:telemetry-rule/telemetry-rule-name
output "observabilityadmin_telemetry_rule" {
  value = provider::arn::observabilityadmin_telemetry_rule("telemetry-rule-name")
}
