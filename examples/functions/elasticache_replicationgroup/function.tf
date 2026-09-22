# arn:aws:elasticache:ap-northeast-1:111111111111:replicationgroup:replication-group-id
output "elasticache_replicationgroup" {
  value = provider::arn::elasticache_replicationgroup("replication-group-id")
}
