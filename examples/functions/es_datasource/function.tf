# arn:aws:opensearch:ap-northeast-1:111111111111:datasource/data-source-name
output "es_datasource" {
  value = provider::arn::es_datasource("data-source-name")
}
