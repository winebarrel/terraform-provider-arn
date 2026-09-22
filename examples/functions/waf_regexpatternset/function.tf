# arn:aws:waf::111111111111:regexpatternset/id
output "waf_regexpatternset" {
  value = provider::arn::waf_regexpatternset("id")
}
