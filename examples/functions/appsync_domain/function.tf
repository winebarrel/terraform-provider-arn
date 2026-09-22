# arn:aws:appsync:ap-northeast-1:111111111111:domainnames/domain-name
output "appsync_domain" {
  value = provider::arn::appsync_domain("domain-name")
}
