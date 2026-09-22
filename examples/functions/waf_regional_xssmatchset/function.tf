# arn:aws:waf-regional:ap-northeast-1:111111111111:xssmatchset/id
output "waf_regional_xssmatchset" {
  value = provider::arn::waf_regional_xssmatchset("id")
}
