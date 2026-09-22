# arn:aws:frauddetector:ap-northeast-1:111111111111:rule/resource-path
output "frauddetector_rule" {
  value = provider::arn::frauddetector_rule("resource-path")
}
