# arn:aws:waf::111111111111:ipset/id
output "waf_ipset" {
  value = provider::arn::waf_ipset("id")
}
