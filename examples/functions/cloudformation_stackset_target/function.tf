# arn:aws:cloudformation:ap-northeast-1:111111111111:stackset-target/stack-set-target
output "cloudformation_stackset_target" {
  value = provider::arn::cloudformation_stackset_target("stack-set-target")
}
