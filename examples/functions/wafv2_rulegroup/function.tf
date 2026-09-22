# arn:aws:wafv2:ap-northeast-1:111111111111:scope/rulegroup/name/id
output "wafv2_rulegroup" {
  value = provider::arn::wafv2_rulegroup("scope", "name", "id")
}
