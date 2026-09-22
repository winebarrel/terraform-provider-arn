# arn:aws:redshift:ap-northeast-1:111111111111:snapshot:cluster-name/snapshot-name
output "redshift_snapshot" {
  value = provider::arn::redshift_snapshot("cluster-name", "snapshot-name")
}
