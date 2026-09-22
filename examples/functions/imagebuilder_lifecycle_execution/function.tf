# arn:aws:imagebuilder:ap-northeast-1:111111111111:lifecycle-execution/lifecycle-execution-id
output "imagebuilder_lifecycle_execution" {
  value = provider::arn::imagebuilder_lifecycle_execution("lifecycle-execution-id")
}
