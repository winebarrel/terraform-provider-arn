# arn:aws:iotsitewise:ap-northeast-1:111111111111:dataset/dataset-id
output "iotsitewise_dataset" {
  value = provider::arn::iotsitewise_dataset("dataset-id")
}
