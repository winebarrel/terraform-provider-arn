# arn:aws:cleanrooms-ml:ap-northeast-1:111111111111:training-dataset/resource-id
output "cleanrooms_ml_trainingdataset" {
  value = provider::arn::cleanrooms_ml_trainingdataset("resource-id")
}
