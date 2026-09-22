# arn:aws:cassandra:ap-northeast-1:111111111111:/keyspace/keyspace-name/table/table-name/stream/stream-label
output "cassandra_stream" {
  value = provider::arn::cassandra_stream("keyspace-name", "table-name", "stream-label")
}
