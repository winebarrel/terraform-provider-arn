# arn:aws:elasticache:ap-northeast-1:111111111111:serverlesscache:serverless-cache-name
output "elasticache_serverlesscache" {
  value = provider::arn::elasticache_serverlesscache("serverless-cache-name")
}
