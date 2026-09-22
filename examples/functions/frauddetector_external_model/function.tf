# arn:aws:frauddetector:ap-northeast-1:111111111111:external-model/resource-path
output "frauddetector_external_model" {
  value = provider::arn::frauddetector_external_model("resource-path")
}
