# arn:aws:iot:ap-northeast-1:111111111111:rule/rule-name
output "iot_rule" {
  value = provider::arn::iot_rule("rule-name")
}
