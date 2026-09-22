# arn:aws:redshift-serverless:ap-northeast-1:111111111111:managed-workgroup/managed-workgroup-id
output "redshift_data_managed_workgroup" {
  value = provider::arn::redshift_data_managed_workgroup("managed-workgroup-id")
}
