# arn:aws:odb:ap-northeast-1:111111111111:db-node/db-node-id
output "odb_db_node" {
  value = provider::arn::odb_db_node("db-node-id")
}
