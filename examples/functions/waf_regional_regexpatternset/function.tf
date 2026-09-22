# arn:aws:waf-regional:ap-northeast-1:111111111111:regexpatternset/id
output "waf_regional_regexpatternset" {
  value = provider::arn::waf_regional_regexpatternset("id")
}
