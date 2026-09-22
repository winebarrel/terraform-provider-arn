# arn:aws:odb:ap-northeast-1:111111111111:exascale-db-storage-vault/exascale-db-storage-vault-id
output "odb_exascale_db_storage_vault" {
  value = provider::arn::odb_exascale_db_storage_vault("exascale-db-storage-vault-id")
}
