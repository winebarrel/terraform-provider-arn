# arn:aws:cleanrooms:ap-northeast-1:111111111111:configuredtable/configured-table-id
output "cleanrooms_configuredtable" {
  value = provider::arn::cleanrooms_configuredtable("configured-table-id")
}
