# arn:aws:workmail:ap-northeast-1:111111111111:organization/resource-id
output "workmail_organization" {
  value = provider::arn::workmail_organization("resource-id")
}
