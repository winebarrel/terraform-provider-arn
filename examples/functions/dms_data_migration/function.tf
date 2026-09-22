# arn:aws:dms:ap-northeast-1:111111111111:data-migration:*
output "dms_data_migration" {
  value = provider::arn::dms_data_migration()
}
