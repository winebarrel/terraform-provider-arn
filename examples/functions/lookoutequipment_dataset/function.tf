# arn:aws:lookoutequipment:ap-northeast-1:111111111111:dataset/dataset-name/dataset-id
output "lookoutequipment_dataset" {
  value = provider::arn::lookoutequipment_dataset("dataset-name", "dataset-id")
}
