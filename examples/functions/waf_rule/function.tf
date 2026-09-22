# arn:aws:waf::111111111111:rule/id
output "waf_rule" {
  value = provider::arn::waf_rule("id")
}
