# arn:aws:elasticache:ap-northeast-1:111111111111:serverlesscachesnapshot:serverless-cache-snapshot-name
output "elasticache_serverlesscachesnapshot" {
  value = provider::arn::elasticache_serverlesscachesnapshot("serverless-cache-snapshot-name")
}
