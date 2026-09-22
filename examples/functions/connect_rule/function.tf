# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/rule/rule-id
output "connect_rule" {
  value = provider::arn::connect_rule("instance-id", "rule-id")
}
