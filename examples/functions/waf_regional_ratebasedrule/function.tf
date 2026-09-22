# arn:aws:waf-regional:ap-northeast-1:111111111111:ratebasedrule/id
output "waf_regional_ratebasedrule" {
  value = provider::arn::waf_regional_ratebasedrule("id")
}
