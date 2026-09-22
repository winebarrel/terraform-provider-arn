# arn:aws:cloudwatch:ap-northeast-1:111111111111:dataset/dataset-id
output "cloudwatch_dataset" {
  value = provider::arn::cloudwatch_dataset("dataset-id")
}
