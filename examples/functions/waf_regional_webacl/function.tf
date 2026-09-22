# arn:aws:waf-regional:ap-northeast-1:111111111111:webacl/id
output "waf_regional_webacl" {
  value = provider::arn::waf_regional_webacl("id")
}
