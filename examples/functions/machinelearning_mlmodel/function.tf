# arn:aws:machinelearning:ap-northeast-1:111111111111:mlmodel/ml-model-id
output "machinelearning_mlmodel" {
  value = provider::arn::machinelearning_mlmodel("ml-model-id")
}
