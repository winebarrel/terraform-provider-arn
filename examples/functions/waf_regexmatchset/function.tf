# arn:aws:waf::111111111111:regexmatch/id
output "waf_regexmatchset" {
  value = provider::arn::waf_regexmatchset("id")
}
