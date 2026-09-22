# arn:aws:redshift-serverless:ap-northeast-1:111111111111:managedvpcendpoint/endpoint-access-id
output "redshift_serverless_endpoint_access" {
  value = provider::arn::redshift_serverless_endpoint_access("endpoint-access-id")
}
