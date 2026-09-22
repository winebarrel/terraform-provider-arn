# arn:aws:cleanrooms-ml:ap-northeast-1:111111111111:configured-model-algorithm/resource-id
output "cleanrooms_ml_configured_model_algorithm" {
  value = provider::arn::cleanrooms_ml_configured_model_algorithm("resource-id")
}
