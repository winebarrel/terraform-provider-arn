# arn:aws:cassandra:ap-northeast-1:111111111111:/keyspace/keyspace-name/
output "cassandra_keyspace" {
  value = provider::arn::cassandra_keyspace("keyspace-name")
}
