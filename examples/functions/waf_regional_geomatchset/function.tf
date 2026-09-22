# arn:aws:waf-regional:ap-northeast-1:111111111111:geomatchset/id
output "waf_regional_geomatchset" {
  value = provider::arn::waf_regional_geomatchset("id")
}
