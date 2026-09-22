# arn:aws:redshift:ap-northeast-1:111111111111:snapshotcopygrant:snapshot-copy-grant-name
output "redshift_snapshotcopygrant" {
  value = provider::arn::redshift_snapshotcopygrant("snapshot-copy-grant-name")
}
