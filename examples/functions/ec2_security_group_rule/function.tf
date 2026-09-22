# arn:aws:ec2:ap-northeast-1:111111111111:security-group-rule/security-group-rule-id
output "ec2_security_group_rule" {
  value = provider::arn::ec2_security_group_rule("security-group-rule-id")
}
