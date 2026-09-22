# arn:aws:translate:ap-northeast-1:111111111111:terminology/resource-name
output "translate_terminology" {
  value = provider::arn::translate_terminology("resource-name")
}
