# arn:aws:frauddetector:ap-northeast-1:111111111111:outcome/resource-path
output "frauddetector_outcome" {
  value = provider::arn::frauddetector_outcome("resource-path")
}
