# arn:aws:redshift:ap-northeast-1:111111111111:dbname:cluster-name/db-name
output "redshift_dbname" {
  value = provider::arn::redshift_dbname("cluster-name", "db-name")
}
