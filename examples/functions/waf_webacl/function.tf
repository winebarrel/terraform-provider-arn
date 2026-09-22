# arn:aws:waf::111111111111:webacl/id
output "waf_webacl" {
  value = provider::arn::waf_webacl("id")
}
