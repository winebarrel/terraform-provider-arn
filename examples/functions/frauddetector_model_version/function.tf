# arn:aws:frauddetector:ap-northeast-1:111111111111:model-version/resource-path
output "frauddetector_model_version" {
  value = provider::arn::frauddetector_model_version("resource-path")
}
