# arn:aws:rds:ap-northeast-1:111111111111:tenant-database:tenant-resource-id
output "rds_tenant_database" {
  value = provider::arn::rds_tenant_database("tenant-resource-id")
}
