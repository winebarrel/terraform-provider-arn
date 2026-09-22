# arn:aws:iotanalytics:ap-northeast-1:111111111111:dataset/dataset-name
output "iotanalytics_dataset" {
  value = provider::arn::iotanalytics_dataset("dataset-name")
}
