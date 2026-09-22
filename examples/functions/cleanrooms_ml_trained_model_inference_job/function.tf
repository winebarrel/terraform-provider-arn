# arn:aws:cleanrooms-ml:ap-northeast-1:111111111111:membership/membership-id/trained-model-inference-job/resource-id
output "cleanrooms_ml_trained_model_inference_job" {
  value = provider::arn::cleanrooms_ml_trained_model_inference_job("membership-id", "resource-id")
}
