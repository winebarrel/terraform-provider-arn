# arn:aws:airflow-serverless:ap-northeast-1:111111111111:workflow/workflow-id
output "airflow_serverless_workflow" {
  value = provider::arn::airflow_serverless_workflow("workflow-id")
}
