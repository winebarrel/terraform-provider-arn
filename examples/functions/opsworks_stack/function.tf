# arn:aws:opsworks:ap-northeast-1:111111111111:stack/stack-id/
output "opsworks_stack" {
  value = provider::arn::opsworks_stack("stack-id")
}
