# arn:aws:sagemaker:ap-northeast-1:111111111111:model/model-name
output "sagemaker_model" {
  value = provider::arn::sagemaker_model("model-name")
}
