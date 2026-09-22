# arn:aws:ivs:ap-northeast-1:111111111111:storage-configuration/resource-id
output "ivs_storage_configuration" {
  value = provider::arn::ivs_storage_configuration("resource-id")
}
