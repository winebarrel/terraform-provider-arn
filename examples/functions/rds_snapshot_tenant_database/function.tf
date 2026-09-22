# arn:aws:rds:ap-northeast-1:111111111111:snapshot-tenant-database:snapshot-name:tenant-resource-id
output "rds_snapshot_tenant_database" {
  value = provider::arn::rds_snapshot_tenant_database("snapshot-name", "tenant-resource-id")
}
