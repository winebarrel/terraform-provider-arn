# arn:aws:redshift-serverless:ap-northeast-1:111111111111:managed-workgroup/managed-workgroup-name
output "redshift_serverless_managed_workgroup" {
  value = provider::arn::redshift_serverless_managed_workgroup("managed-workgroup-name")
}
