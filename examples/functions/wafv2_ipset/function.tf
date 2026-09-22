# arn:aws:wafv2:ap-northeast-1:111111111111:scope/ipset/name/id
output "wafv2_ipset" {
  value = provider::arn::wafv2_ipset("scope", "name", "id")
}
