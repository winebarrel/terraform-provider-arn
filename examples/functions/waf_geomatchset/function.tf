# arn:aws:waf::111111111111:geomatchset/id
output "waf_geomatchset" {
  value = provider::arn::waf_geomatchset("id")
}
