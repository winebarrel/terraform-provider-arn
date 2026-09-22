# arn:aws:airflow:ap-northeast-1:111111111111:role/environment-name/role-name
output "airflow_rbac_role" {
  value = provider::arn::airflow_rbac_role("environment-name", "role-name")
}
