# arn:aws:cloudformation:ap-northeast-1:111111111111:type/resource/type
output "cloudformation_type" {
  value = provider::arn::cloudformation_type("type")
}
