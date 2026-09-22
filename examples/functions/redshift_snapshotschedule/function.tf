# arn:aws:redshift:ap-northeast-1:111111111111:snapshotschedule:schedule-identifier
output "redshift_snapshotschedule" {
  value = provider::arn::redshift_snapshotschedule("schedule-identifier")
}
