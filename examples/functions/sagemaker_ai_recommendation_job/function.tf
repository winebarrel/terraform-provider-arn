# arn:aws:sagemaker:ap-northeast-1:111111111111:ai-recommendation-job/ai-recommendation-job-name
output "sagemaker_ai_recommendation_job" {
  value = provider::arn::sagemaker_ai_recommendation_job("ai-recommendation-job-name")
}
