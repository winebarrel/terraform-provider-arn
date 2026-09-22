# arn:aws:redshift:ap-northeast-1:111111111111:dbgroup:cluster-name/db-group
output "redshift_dbgroup" {
  value = provider::arn::redshift_dbgroup("cluster-name", "db-group")
}
