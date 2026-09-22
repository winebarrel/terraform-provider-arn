# arn:aws:waf::111111111111:ratebasedrule/id
output "waf_ratebasedrule" {
  value = provider::arn::waf_ratebasedrule("id")
}
