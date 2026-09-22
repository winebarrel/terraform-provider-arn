# arn:aws:frauddetector:ap-northeast-1:111111111111:variable/resource-path
output "frauddetector_variable" {
  value = provider::arn::frauddetector_variable("resource-path")
}
