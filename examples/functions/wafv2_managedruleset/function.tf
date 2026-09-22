# arn:aws:wafv2:ap-northeast-1:111111111111:scope/managedruleset/name/id
output "wafv2_managedruleset" {
  value = provider::arn::wafv2_managedruleset("scope", "name", "id")
}
