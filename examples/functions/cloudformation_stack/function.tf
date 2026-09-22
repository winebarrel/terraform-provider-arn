# arn:aws:cloudformation:ap-northeast-1:111111111111:stack/stack-name/id
output "cloudformation_stack" {
  value = provider::arn::cloudformation_stack("stack-name", "id")
}
