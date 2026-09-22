# arn:aws:workdocs:ap-northeast-1:111111111111:organization/resource-id
output "workdocs_organization" {
  value = provider::arn::workdocs_organization("resource-id")
}
