# arn:aws:sagemaker:ap-northeast-1:111111111111:notebook-instance/notebook-instance-name
output "sagemaker_notebook_instance" {
  value = provider::arn::sagemaker_notebook_instance("notebook-instance-name")
}
