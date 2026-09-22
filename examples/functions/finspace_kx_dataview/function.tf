# arn:aws:finspace:ap-northeast-1:111111111111:kxEnvironment/environment-id/kxDatabase/kx-database/kxDataview/kx-dataview
output "finspace_kx_dataview" {
  value = provider::arn::finspace_kx_dataview("environment-id", "kx-database", "kx-dataview")
}
