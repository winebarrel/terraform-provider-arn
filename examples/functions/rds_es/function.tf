# arn:aws:rds:ap-northeast-1:111111111111:es:subscription-name
output "rds_es" {
  value = provider::arn::rds_es("subscription-name")
}
