# arn:aws:cloudwatch:ap-northeast-1:111111111111:alarm-mute-rule:alarm-mute-rule-name
output "cloudwatch_alarm_mute_rule" {
  value = provider::arn::cloudwatch_alarm_mute_rule("alarm-mute-rule-name")
}
