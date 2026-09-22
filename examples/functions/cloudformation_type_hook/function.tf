# arn:aws:cloudformation:ap-northeast-1:111111111111:type/hook/type
output "cloudformation_type_hook" {
  value = provider::arn::cloudformation_type_hook("type")
}
