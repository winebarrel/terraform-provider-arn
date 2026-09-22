# arn:aws:frauddetector:ap-northeast-1:111111111111:detector/resource-path
output "frauddetector_detector" {
  value = provider::arn::frauddetector_detector("resource-path")
}
