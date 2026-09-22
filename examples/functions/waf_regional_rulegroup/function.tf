# arn:aws:waf-regional:ap-northeast-1:111111111111:rulegroup/id
output "waf_regional_rulegroup" {
  value = provider::arn::waf_regional_rulegroup("id")
}
