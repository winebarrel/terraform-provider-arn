# arn:aws:elasticache:ap-northeast-1:111111111111:cluster:cache-cluster-id
output "elasticache_cluster" {
  value = provider::arn::elasticache_cluster("cache-cluster-id")
}
