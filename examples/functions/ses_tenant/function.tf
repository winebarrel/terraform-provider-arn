# arn:aws:ses:ap-northeast-1:111111111111:tenant/tenant-name/tenant-id
output "ses_tenant" {
  value = provider::arn::ses_tenant("tenant-name", "tenant-id")
}
