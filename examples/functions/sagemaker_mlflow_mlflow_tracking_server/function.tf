# arn:aws:sagemaker:ap-northeast-1:111111111111:mlflow-tracking-server/mlflow-tracking-server-name
output "sagemaker_mlflow_mlflow_tracking_server" {
  value = provider::arn::sagemaker_mlflow_mlflow_tracking_server("mlflow-tracking-server-name")
}
