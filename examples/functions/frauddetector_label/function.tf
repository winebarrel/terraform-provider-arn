# arn:aws:frauddetector:ap-northeast-1:111111111111:label/resource-path
output "frauddetector_label" {
  value = provider::arn::frauddetector_label("resource-path")
}
