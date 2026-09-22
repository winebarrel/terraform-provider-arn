# arn:aws:frauddetector:ap-northeast-1:111111111111:model/resource-path
output "frauddetector_model" {
  value = provider::arn::frauddetector_model("resource-path")
}
