# arn:aws:frauddetector:ap-northeast-1:111111111111:batch-prediction/resource-path
output "frauddetector_batch_prediction" {
  value = provider::arn::frauddetector_batch_prediction("resource-path")
}
