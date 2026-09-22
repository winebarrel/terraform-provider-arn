# arn:aws:waf-regional:ap-northeast-1:111111111111:bytematchset/id
output "waf_regional_bytematchset" {
  value = provider::arn::waf_regional_bytematchset("id")
}
