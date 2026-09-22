# arn:aws:cloudformation:ap-northeast-1:111111111111:stackset/stack-set-name:id
output "cloudformation_stackset" {
  value = provider::arn::cloudformation_stackset("stack-set-name", "id")
}
