# arn:aws:sagemaker:ap-northeast-1:111111111111:notebook-instance-lifecycle-config/notebook-instance-lifecycle-config-name
output "sagemaker_notebook_instance_lifecycle_config" {
  value = provider::arn::sagemaker_notebook_instance_lifecycle_config("notebook-instance-lifecycle-config-name")
}
