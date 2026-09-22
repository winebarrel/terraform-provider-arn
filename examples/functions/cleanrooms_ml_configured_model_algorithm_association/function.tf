# arn:aws:cleanrooms-ml:ap-northeast-1:111111111111:membership/membership-id/configured-model-algorithm-association/resource-id
output "cleanrooms_ml_configured_model_algorithm_association" {
  value = provider::arn::cleanrooms_ml_configured_model_algorithm_association("membership-id", "resource-id")
}
