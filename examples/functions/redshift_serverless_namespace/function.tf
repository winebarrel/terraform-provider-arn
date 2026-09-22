# arn:aws:redshift-serverless:ap-northeast-1:111111111111:namespace/namespace-id
output "redshift_serverless_namespace" {
  value = provider::arn::redshift_serverless_namespace("namespace-id")
}
