# arn:aws:waf-regional:ap-northeast-1:111111111111:ipset/id
output "waf_regional_ipset" {
  value = provider::arn::waf_regional_ipset("id")
}
