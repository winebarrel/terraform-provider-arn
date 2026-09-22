# arn:aws:elasticache:ap-northeast-1:111111111111:snapshot:snapshot-name
output "elasticache_snapshot" {
  value = provider::arn::elasticache_snapshot("snapshot-name")
}
