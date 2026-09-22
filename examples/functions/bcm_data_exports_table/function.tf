# arn:aws:bcm-data-exports:ap-northeast-1:111111111111:table/identifier
output "bcm_data_exports_table" {
  value = provider::arn::bcm_data_exports_table("identifier")
}
