# arn:aws:observabilityadmin:ap-northeast-1:111111111111:dataset-integration/dataset-integration-identifier
output "observabilityadmin_dataset_integration" {
  value = provider::arn::observabilityadmin_dataset_integration("dataset-integration-identifier")
}
