# arn:aws:sagemaker:ap-northeast-1:111111111111:ai-benchmark-job/ai-benchmark-job-name
output "sagemaker_ai_benchmark_job" {
  value = provider::arn::sagemaker_ai_benchmark_job("ai-benchmark-job-name")
}
