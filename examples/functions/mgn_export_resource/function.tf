# arn:aws:mgn:ap-northeast-1:111111111111:export/export-id
output "mgn_export_resource" {
  value = provider::arn::mgn_export_resource("export-id")
}
