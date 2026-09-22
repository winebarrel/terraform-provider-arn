# arn:aws:ssm:ap-northeast-1:111111111111:resource-data-sync/sync-name
output "ssm_resourcedatasync" {
  value = provider::arn::ssm_resourcedatasync("sync-name")
}
