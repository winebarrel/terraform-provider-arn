# arn:aws:timestream-influxdb:ap-northeast-1:111111111111:db-instance/db-instance-identifier
output "timestream_influxdb_db_instance" {
  value = provider::arn::timestream_influxdb_db_instance("db-instance-identifier")
}
