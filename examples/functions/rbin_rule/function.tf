# arn:aws:rbin:ap-northeast-1:111111111111:rule/resource-name
output "rbin_rule" {
  value = provider::arn::rbin_rule("resource-name")
}
