# arn:aws:wafv2:ap-northeast-1:111111111111:scope/regexpatternset/name/id
output "wafv2_regexpatternset" {
  value = provider::arn::wafv2_regexpatternset("scope", "name", "id")
}
