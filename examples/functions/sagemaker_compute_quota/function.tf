# arn:aws:sagemaker:ap-northeast-1:111111111111:compute-quota/compute-quota-id
output "sagemaker_compute_quota" {
  value = provider::arn::sagemaker_compute_quota("compute-quota-id")
}
