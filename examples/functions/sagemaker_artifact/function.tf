# arn:aws:sagemaker:ap-northeast-1:111111111111:artifact/hash-of-artifact-source
output "sagemaker_artifact" {
  value = provider::arn::sagemaker_artifact("hash-of-artifact-source")
}
