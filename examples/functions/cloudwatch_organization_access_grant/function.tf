# arn:aws:cloudwatch:ap-northeast-1:111111111111:organization-access-grant/grant-id
output "cloudwatch_organization_access_grant" {
  value = provider::arn::cloudwatch_organization_access_grant("grant-id")
}
