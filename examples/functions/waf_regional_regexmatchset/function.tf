# arn:aws:waf-regional:ap-northeast-1:111111111111:regexmatch/id
output "waf_regional_regexmatchset" {
  value = provider::arn::waf_regional_regexmatchset("id")
}
