# arn:aws:waf::111111111111:xssmatchset/id
output "waf_xssmatchset" {
  value = provider::arn::waf_xssmatchset("id")
}
