# arn:aws:timestream-influxdb:ap-northeast-1:111111111111:db-cluster/db-cluster-id
output "timestream_influxdb_db_cluster" {
  value = provider::arn::timestream_influxdb_db_cluster("db-cluster-id")
}
