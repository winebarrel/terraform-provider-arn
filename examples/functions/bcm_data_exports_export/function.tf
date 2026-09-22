# arn:aws:bcm-data-exports:ap-northeast-1:111111111111:export/identifier
output "bcm_data_exports_export" {
  value = provider::arn::bcm_data_exports_export("identifier")
}
