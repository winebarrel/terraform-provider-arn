# arn:aws:cloudwatch:ap-northeast-1:111111111111:organization-domain/domain-id
output "cloudwatch_organization_domain" {
  value = provider::arn::cloudwatch_organization_domain("domain-id")
}
