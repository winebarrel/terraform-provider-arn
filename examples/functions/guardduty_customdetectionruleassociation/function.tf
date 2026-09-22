# arn:aws:guardduty:ap-northeast-1:111111111111:detection-rule/custom/rule-id/association/association-id
output "guardduty_customdetectionruleassociation" {
  value = provider::arn::guardduty_customdetectionruleassociation("rule-id", "association-id")
}
