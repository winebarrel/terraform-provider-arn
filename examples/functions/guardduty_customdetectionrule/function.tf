# arn:aws:guardduty::aws:detection-rule/custom/rule-id
output "guardduty_customdetectionrule" {
  value = provider::arn::guardduty_customdetectionrule("rule-id")
}
