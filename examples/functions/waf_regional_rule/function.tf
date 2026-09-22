# arn:aws:waf-regional:ap-northeast-1:111111111111:rule/id
output "waf_regional_rule" {
  value = provider::arn::waf_regional_rule("id")
}
