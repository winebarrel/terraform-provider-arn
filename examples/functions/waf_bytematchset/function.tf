# arn:aws:waf::111111111111:bytematchset/id
output "waf_bytematchset" {
  value = provider::arn::waf_bytematchset("id")
}
