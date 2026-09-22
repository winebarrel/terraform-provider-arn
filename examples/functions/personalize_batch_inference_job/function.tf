# arn:aws:personalize:ap-northeast-1:111111111111:batch-inference-job/resource-id
output "personalize_batch_inference_job" {
  value = provider::arn::personalize_batch_inference_job("resource-id")
}
