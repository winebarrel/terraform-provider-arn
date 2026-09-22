# arn:aws:sagemaker:ap-northeast-1:111111111111:inference-recommendations-job/inference-recommendations-job-name
output "sagemaker_inference_recommendations_job" {
  value = provider::arn::sagemaker_inference_recommendations_job("inference-recommendations-job-name")
}
