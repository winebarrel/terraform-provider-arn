# arn:aws:sdb:ap-northeast-1:111111111111:domain/domain-name
output "sdb_domain" {
  value = provider::arn::sdb_domain("domain-name")
}
