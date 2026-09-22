# arn:aws:waf::111111111111:sizeconstraintset/id
output "waf_sizeconstraintset" {
  value = provider::arn::waf_sizeconstraintset("id")
}
