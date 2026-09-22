# arn:aws:elasticache:ap-northeast-1:111111111111:reserved-instance:reserved-cache-node-id
output "elasticache_reserved_instance" {
  value = provider::arn::elasticache_reserved_instance("reserved-cache-node-id")
}
