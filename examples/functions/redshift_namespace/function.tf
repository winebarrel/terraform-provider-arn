# arn:aws:redshift:ap-northeast-1:111111111111:namespace:cluster-namespace
output "redshift_namespace" {
  value = provider::arn::redshift_namespace("cluster-namespace")
}
