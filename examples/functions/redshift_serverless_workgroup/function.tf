# arn:aws:redshift-serverless:ap-northeast-1:111111111111:workgroup/workgroup-id
output "redshift_serverless_workgroup" {
  value = provider::arn::redshift_serverless_workgroup("workgroup-id")
}
