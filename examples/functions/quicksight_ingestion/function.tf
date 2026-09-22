# arn:aws:quicksight:ap-northeast-1:111111111111:dataset/dataset-id/ingestion/resource-id
output "quicksight_ingestion" {
  value = provider::arn::quicksight_ingestion("dataset-id", "resource-id")
}
