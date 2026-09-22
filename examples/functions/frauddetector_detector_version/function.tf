# arn:aws:frauddetector:ap-northeast-1:111111111111:detector-version/resource-path
output "frauddetector_detector_version" {
  value = provider::arn::frauddetector_detector_version("resource-path")
}
