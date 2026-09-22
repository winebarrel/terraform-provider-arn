# arn:aws:frauddetector:ap-northeast-1:111111111111:batch-import/resource-path
output "frauddetector_batch_import" {
  value = provider::arn::frauddetector_batch_import("resource-path")
}
