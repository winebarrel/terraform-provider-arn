# arn:aws:cleanrooms-ml:ap-northeast-1:111111111111:membership/membership-id/trained-model/resource-id
output "cleanrooms_ml_trained_model" {
  value = provider::arn::cleanrooms_ml_trained_model("membership-id", "resource-id")
}
