# arn:aws:appstream:ap-northeast-1:111111111111:stack/stack-name
output "appstream_stack" {
  value = provider::arn::appstream_stack("stack-name")
}
