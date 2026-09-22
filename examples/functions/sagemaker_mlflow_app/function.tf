# arn:aws:sagemaker:ap-northeast-1:111111111111:mlflow-app/m-lflow-app-id
output "sagemaker_mlflow_app" {
  value = provider::arn::sagemaker_mlflow_app("m-lflow-app-id")
}
