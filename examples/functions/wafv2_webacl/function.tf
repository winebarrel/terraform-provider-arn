# arn:aws:wafv2:ap-northeast-1:111111111111:scope/webacl/name/id
output "wafv2_webacl" {
  value = provider::arn::wafv2_webacl("scope", "name", "id")
}
