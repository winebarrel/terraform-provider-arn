# arn:aws:sagemaker:ap-northeast-1:111111111111:feature-group/feature-group-name
output "sagemaker_feature_group" {
  value = provider::arn::sagemaker_feature_group("feature-group-name")
}
