# arn:aws:xray:ap-northeast-1:111111111111:sampling-rule/sampling-rule-name
output "xray_sampling_rule" {
  value = provider::arn::xray_sampling_rule("sampling-rule-name")
}
