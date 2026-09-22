# arn:aws:ses:ap-northeast-1:111111111111:mailmanager-rule-set/rule-set-id
output "ses_mailmanager_rule_set" {
  value = provider::arn::ses_mailmanager_rule_set("rule-set-id")
}
