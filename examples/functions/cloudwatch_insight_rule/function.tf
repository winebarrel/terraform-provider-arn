# arn:aws:cloudwatch:ap-northeast-1:111111111111:insight-rule/insight-rule-name
output "cloudwatch_insight_rule" {
  value = provider::arn::cloudwatch_insight_rule("insight-rule-name")
}
