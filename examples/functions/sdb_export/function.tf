# arn:aws:sdb:ap-northeast-1:111111111111:domain/domain-name/export/export-uuid
output "sdb_export" {
  value = provider::arn::sdb_export("domain-name", "export-uuid")
}
