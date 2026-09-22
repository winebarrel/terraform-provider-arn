# arn:aws:sagemaker:ap-northeast-1:111111111111:reserved-capacity/random-string
output "sagemaker_reserved_capacity" {
  value = provider::arn::sagemaker_reserved_capacity("random-string")
}
