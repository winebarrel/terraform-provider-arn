# arn:aws:sagemaker:ap-northeast-1:111111111111:model-card/model-card-name
output "sagemaker_model_card" {
  value = provider::arn::sagemaker_model_card("model-card-name")
}
