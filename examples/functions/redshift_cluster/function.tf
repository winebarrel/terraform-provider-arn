# arn:aws:redshift:ap-northeast-1:111111111111:cluster:cluster-name
output "redshift_cluster" {
  value = provider::arn::redshift_cluster("cluster-name")
}
