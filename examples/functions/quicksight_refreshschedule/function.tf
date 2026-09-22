# arn:aws:quicksight:ap-northeast-1:111111111111:dataset/dataset-id/refresh-schedule/resource-id
output "quicksight_refreshschedule" {
  value = provider::arn::quicksight_refreshschedule("dataset-id", "resource-id")
}
