# arn:aws:mgn:ap-northeast-1:111111111111:import/import-id
output "mgn_import_resource" {
  value = provider::arn::mgn_import_resource("import-id")
}
