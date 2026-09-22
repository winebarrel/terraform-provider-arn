# arn:aws:sagemaker:ap-northeast-1:111111111111:shared-model/shared-model-id
output "sagemaker_shared_model" {
  value = provider::arn::sagemaker_shared_model("shared-model-id")
}
