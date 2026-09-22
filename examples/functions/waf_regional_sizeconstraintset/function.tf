# arn:aws:waf-regional:ap-northeast-1:111111111111:sizeconstraintset/id
output "waf_regional_sizeconstraintset" {
  value = provider::arn::waf_regional_sizeconstraintset("id")
}
