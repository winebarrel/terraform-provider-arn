# arn:aws:cloudwatch:ap-northeast-1:111111111111:domain/domain-id
output "cloudwatch_domain" {
  value = provider::arn::cloudwatch_domain("domain-id")
}
