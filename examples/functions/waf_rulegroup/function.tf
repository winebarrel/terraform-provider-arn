# arn:aws:waf::111111111111:rulegroup/id
output "waf_rulegroup" {
  value = provider::arn::waf_rulegroup("id")
}
