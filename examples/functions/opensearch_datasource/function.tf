# arn:aws:opensearch:ap-northeast-1:111111111111:datasource/data-source-name
output "opensearch_datasource" {
  value = provider::arn::opensearch_datasource("data-source-name")
}
