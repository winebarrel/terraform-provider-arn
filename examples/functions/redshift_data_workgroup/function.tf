# arn:aws:redshift-serverless:ap-northeast-1:111111111111:workgroup/workgroup-id
output "redshift_data_workgroup" {
  value = provider::arn::redshift_data_workgroup("workgroup-id")
}
