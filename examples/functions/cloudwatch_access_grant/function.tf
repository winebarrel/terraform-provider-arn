# arn:aws:cloudwatch:ap-northeast-1:111111111111:access-grant/grant-id
output "cloudwatch_access_grant" {
  value = provider::arn::cloudwatch_access_grant("grant-id")
}
