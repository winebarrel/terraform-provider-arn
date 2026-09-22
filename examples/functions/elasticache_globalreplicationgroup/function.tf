# arn:aws:elasticache::111111111111:globalreplicationgroup:global-replication-group-id
output "elasticache_globalreplicationgroup" {
  value = provider::arn::elasticache_globalreplicationgroup("global-replication-group-id")
}
