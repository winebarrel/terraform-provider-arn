# arn:aws:redshift:ap-northeast-1:111111111111:dbuser:cluster-name/db-user
output "redshift_dbuser" {
  value = provider::arn::redshift_dbuser("cluster-name", "db-user")
}
