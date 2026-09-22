# arn:aws:sagemaker:ap-northeast-1:111111111111:ai-workload-config/ai-workload-config-name
output "sagemaker_ai_workload_config" {
  value = provider::arn::sagemaker_ai_workload_config("ai-workload-config-name")
}
