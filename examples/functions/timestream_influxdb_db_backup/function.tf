# arn:aws:timestream-influxdb:ap-northeast-1:111111111111:db-backup/db-backup-id
output "timestream_influxdb_db_backup" {
  value = provider::arn::timestream_influxdb_db_backup("db-backup-id")
}
