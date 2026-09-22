# arn:aws:timestream-influxdb:ap-northeast-1:111111111111:db-parameter-group/db-parameter-group-identifier
output "timestream_influxdb_db_parameter_group" {
  value = provider::arn::timestream_influxdb_db_parameter_group("db-parameter-group-identifier")
}
