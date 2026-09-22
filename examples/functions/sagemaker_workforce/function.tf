# arn:aws:sagemaker:ap-northeast-1:111111111111:workforce/workforce-name
output "sagemaker_workforce" {
  value = provider::arn::sagemaker_workforce("workforce-name")
}
