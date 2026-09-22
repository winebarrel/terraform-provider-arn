# arn:aws:cassandra:ap-northeast-1:111111111111:/keyspace/keyspace-name/table/table-name
output "cassandra_table" {
  value = provider::arn::cassandra_table("keyspace-name", "table-name")
}
