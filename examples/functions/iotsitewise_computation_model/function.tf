# arn:aws:iotsitewise:ap-northeast-1:111111111111:computation-model/computation-model-id
output "iotsitewise_computation_model" {
  value = provider::arn::iotsitewise_computation_model("computation-model-id")
}
