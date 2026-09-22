# arn:aws:machinelearning:ap-northeast-1:111111111111:batchprediction/batch-prediction-id
output "machinelearning_batchprediction" {
  value = provider::arn::machinelearning_batchprediction("batch-prediction-id")
}
