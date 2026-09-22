# arn:aws:airflow:ap-northeast-1:111111111111:environment/environment-name
output "airflow_environment" {
  value = provider::arn::airflow_environment("environment-name")
}
